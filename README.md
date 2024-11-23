# Computer Graphics Project

Этот проект предназначен для экспериментов и изучения различных методов обработки изображений и алгоритмов компьютерной графики с использованием языка программирования Go и библиотеки GoCV (обертка для OpenCV).

## Структура директорий

- `cmd/`  
  Содержит исходный код для выполнения лабораторных работ по компьютерной графике.

  - `cmd/1_lab/`
    - `main.go`: Код для первой лабораторной работы, в которой проводятся операции над изображениями.

  - `cmd/2_lab/`
    - `main.go`: Код для второй лабораторной работы, включая алгоритмы обработки изображений, такие как дизеринг (Floyd-Steinberg) и преобразование изображения в черно-белое.

  - `cmd/3_lab/`
      - `main.go`: Код для второй лабораторной работы. Реализован алгоритм Брезенхема, алгоритм построения многоугольника по точкам, алгоритмы проверки на выпуклость/не выпуклость и присутствие самопересечений, а также метод Even-Odd и Non-Zero Winding для заполнения полигонов.

- `static/`  
  Папка для хранения изображений, используемых в лабораторных работах.


## Установка

Для работы с этим проектом необходимо установить библиотеку GoCV. Пожалуйста, ознакомьтесь с документацией по установке [здесь](https://gocv.io/).

## Запуск

Для запуска отдельных лабораторных работ можно использовать команды:

```bash
# Для первой лабораторной работы
go run cmd/1_lab/main.go

# Для второй лабораторной работы
go run cmd/2_lab/main.go

# Для третьей лабораторной работы
go run cmd/3_lab/main.go
