# Time Tracker
Controla de una forma simple el tiempo que dedicas a tus tareas mediante pomodoros.

Esta herramienta está pensada para llevar un registro de tiempo, organizado en proyectos y tareas, de esta forma podrás saber cuánto tiempo exactamente dedicas a cada cosa.

## Uso

```bash
$ go get
$ cd tracker/
$ go build
$ ./tracker help
Time tracker ejecuta pomodoros sobre tareas y registra el tiempo en cada una de ellas.

Usage:
  tracker [command]

Available Commands:
  add         Añade Proyectos o tareas a proyectos
  checkout    Cambia el contexto a una tarea determinada
  delete      Borra una tarea o un proyecto
  help        Help about any command
  list        Lista proyectos y tareas
  status      Resumen del estado actual
  track       Inicia un pomodoro y registra el tiempo en una tarea determinada

Flags:
  -h, --help   help for tracker

Use "tracker [command] --help" for more information about a command.

```


## Licencia
MIT 
