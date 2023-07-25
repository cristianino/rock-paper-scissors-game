"use strics"

let imgArray = [
    "/public/img/hand_rock.png",
    "/public/img/hand_paper.png",
    "/public/img/hand_scissors.png",
]

function choose(x) {

    if (x == 0) {
        document.getElementById("player_choice").innerHTML = "El jugador eligió PIEDRA."
    } else if (x == 1) {
        document.getElementById("player_choice").innerHTML = "El jugador eligió PAPEL."
    } else {
        document.getElementById("player_choice").innerHTML = "El jugador eligió TIJERAS."
    }
    fetch("/play?c=" + x).then(res => res.json())
        .then(data => {
            document.getElementById("player_score").innerHTML = data.player_score
            document.getElementById("computer_score").innerHTML = data.computer_score

            document.getElementById("computer_choice").innerHTML = data.computer_choice
            document.getElementById("round_result").innerHTML = data.round_result
            document.getElementById("round_message").innerHTML = data.message

            var imgElement = document.getElementById("img_computer")

            imgElement.src = imgArray[data.computer_choice_int]
        })
        .catch(e => {
            console.error(e);
        })
}

