# RPS WEB


## Determinación del ganador por relación circular

Este código en Go define tres constantes numéricas (ROCK, PAPER y SCISSORS) que representan las opciones del juego piedra, papel o tijeras. Cada una de estas constantes está asociada a un número entero: 0 para ROCK (piedra), 1 para PAPER (papel) y 2 para SCISSORS (tijeras).

Para determinar quién gana en este juego, se utiliza una regla simple basada en aritmética modular (también conocida como "cálculo módulo" o "resto de la división"). La regla es la siguiente:

Piedra (ROCK) vence a Tijeras (SCISSORS) porque (SCISSORS + 1) % 3 = 0, y el resto de dividir 3 por 1 es 0.
Papel (PAPER) vence a Piedra (ROCK) porque (ROCK + 1) % 3 = 1, y el resto de dividir 3 por 2 es 1.
Tijeras (SCISSORS) vence a Papel (PAPER) porque (PAPER + 1) % 3 = 2, y el resto de dividir 3 por 3 es 2.
La clave para entender cómo funciona este algoritmo está en cómo se relacionan los valores numéricos de las opciones y cómo se aplican las operaciones de aritmética modular para determinar el ganador.

Ahora, veamos cómo se interpreta este algoritmo de forma gráfica con algunos ejemplos:

Piedra vs. Tijeras:

Elección del jugador 1 (P1): Piedra (ROCK) => Representado por el número 0.
Elección del jugador 2 (P2): Tijeras (SCISSORS) => Representado por el número 2.
Para determinar el ganador:

(P1 + 1) % 3 = (0 + 1) % 3 = 1.
(P2 + 1) % 3 = (2 + 1) % 3 = 0.
Como el resultado de (P1 + 1) % 3 = 1 y (P2 + 1) % 3 = 0, el Jugador 1 (P1) gana porque Piedra vence a Tijeras.

Papel vs. Piedra:

Elección del jugador 1 (P1): Papel (PAPER) => Representado por el número 1.
Elección del jugador 2 (P2): Piedra (ROCK) => Representado por el número 0.
Para determinar el ganador:

(P1 + 1) % 3 = (1 + 1) % 3 = 2.
(P2 + 1) % 3 = (0 + 1) % 3 = 1.
Como el resultado de (P1 + 1) % 3 = 2 y (P2 + 1) % 3 = 1, el Jugador 1 (P1) gana porque Papel vence a Piedra.

Tijeras vs. Papel:

Elección del jugador 1 (P1): Tijeras (SCISSORS) => Representado por el número 2.
Elección del jugador 2 (P2): Papel (PAPER) => Representado por el número 1.
Para determinar el ganador:

(P1 + 1) % 3 = (2 + 1) % 3 = 0.
(P2 + 1) % 3 = (1 + 1) % 3 = 2.
Como el resultado de (P1 + 1) % 3 = 0 y (P2 + 1) % 3 = 2, el Jugador 2 (P2) gana porque Tijeras vence a Papel.