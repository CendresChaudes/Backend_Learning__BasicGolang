# Подключается из ~/.zshrc.
# В папке урока make next и make prev выполняют cd в текущем терминале.

if [[ -n ${ZSH_VERSION:-} && ${ZSH_EVAL_CONTEXT:-} != *:file* ]]; then
  echo "Подключи файл в текущий терминал: source nav.zsh"
  exit 1
fi

_lesson_cd() {
  local key="$1"
  local empty_msg="$2"
  local line dir

  if [[ ! -f Makefile ]]; then
    echo "В этой папке нет Makefile урока."
    return 1
  fi

  line=$(grep -E "^${key} := " Makefile | head -n 1)
  dir="${line#"${key} := "}"
  dir="${dir%"${dir##*[![:space:]]}"}"

  if [[ -z "$dir" ]]; then
    echo "$empty_msg"
    return 1
  fi

  if [[ ! -d "$dir" ]]; then
    echo "Папка не найдена: $dir"
    return 1
  fi

  cd "$dir" || return 1
  pwd
}

next() {
  _lesson_cd NEXT "Это последний урок."
}

prev() {
  _lesson_cd PREV "Это первый урок."
}

make() {
  case "$1" in
    next|prev)
      local key
      key=${(U)1}
      if [[ -f Makefile ]] && grep -q -E "^${key} := " Makefile; then
        _lesson_cd "$key" "$([[ "$key" == NEXT ]] && echo "Это последний урок." || echo "Это первый урок.")"
        return $?
      fi
      ;;
  esac
  command make "$@"
}
