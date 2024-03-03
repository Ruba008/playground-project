<script setup>
import { ref } from "vue";
const player = ref("Player 1");

const square11 = ref(null);
const square12 = ref(null);
const square13 = ref(null);
const square21 = ref(null);
const square22 = ref(null);
const square23 = ref(null);
const square31 = ref(null);
const square32 = ref(null);
const square33 = ref(null);

// Defining if is the player's turn 1 or 2
// true: player 1
// false: player 2
const playerBool = ref(true);

// Function that change the player's turn.
function changePlayer(){
  playerBool.value = playerBool.value ? false : true;
  player.value = playerBool.value ? "Player 1" : "Player 2";
}

function setPosition(position){
  
   // Consuming API 
  fetch("http://localhost:8081/v1/hash/position", {
    method: "POST",
    body: JSON.stringify({
      player: !playerBool.value,
      position: position,
    }),
  }).then((res) => console.log(res.json()));
}

function getPosition(position) {

  // Defining X or O 
  switch (position.toString()) {
    case [1, 1].toString():
      console.log("aki");
      if (square11.value == null) {
        changePlayer()
        setPosition(position)
        square11.value = !playerBool.value;
      }
      break;
    case [1, 2].toString():
      if (square12.value == null) {
        changePlayer()
        setPosition(position)
        square12.value = !playerBool.value;
      }
      break;
    case [1, 3].toString():
      if (square13.value == null) {
        changePlayer()
        setPosition(position)
        square13.value = !playerBool.value;
      }
      break;
    case [2, 1].toString():
      if (square21.value == null) {
        changePlayer()
        setPosition(position)
        square21.value = !playerBool.value;
      }
      break;
    case [2, 2].toString():
      if (square22.value == null) {
        changePlayer()
        setPosition(position)
        square22.value = !playerBool.value;
      }
      break;
    case [2, 3].toString():
      if (square23.value == null) {
        changePlayer()
        setPosition(position)
        square23.value = !playerBool.value;
      }
      break;
    case [3, 1].toString():
      if (square31.value == null) {
        changePlayer()
        setPosition(position)
        square31.value = !playerBool.value;
      }
      break;
    case [3, 2].toString():
      if (square32.value == null) {
        changePlayer()
        setPosition(position)
        square32.value = !playerBool.value;
      }
      break;
    case [3, 3].toString():
      if (square33.value == null) {
        changePlayer()
        setPosition(position)
        square33.value = !playerBool.value;
      }
      break;
  }

}
</script>

<template>
  <div class="flex w-full justify-center flex-col items-center">
    <div class="grid grid-cols-3 text-center w-60 h-60 mt-10">
      <div
        class="flex w-[80px] h-[80px] cursor-pointer items-center justify-center border-r-4 border-b-4 border-black"
        @click="getPosition([1, 1])"
      >
        <div class="flex m-2 justify-center items-center h-full">
          <img v-if="square11 == true" src="../assets/circlePlay.svg" alt="" />
          <img
            class="w-[48px]"
            v-else-if="square11 == false"
            src="../assets/crossPlay.svg"
            alt=""
          />
        </div>
      </div>
      <div
        class="cursor-pointer w-[80px] h-[80px] flex items-center justify-center border-r-4 border-b-4 border-black"
        @click="getPosition([1, 2])"
      >
        <div class="flex m-2 justify-center items-center h-full">
          <img v-if="square12 == true" src="../assets/circlePlay.svg" alt="" />
          <img
            class="w-[48px]"
            v-else-if="square12 == false"
            src="../assets/crossPlay.svg"
            alt=""
          />
        </div>
      </div>
      <div
        class="cursor-pointer w-[80px] h-[80px] flex items-center justify-center border-black border-b-4"
        @click="getPosition([1, 3])"
      >
        <div class="flex m-2 justify-center items-center h-full">
          <img v-if="square13 == true" src="../assets/circlePlay.svg" alt="" />
          <img
            class="w-[48px]"
            v-else-if="square13 == false"
            src="../assets/crossPlay.svg"
            alt=""
          />
        </div>
      </div>
      <div
        class="cursor-pointer w-[80px] h-[80px] flex items-center justify-center border-black border-b-4 border-r-4"
        @click="getPosition([2, 1])"
      >
        <div class="flex m-2 justify-center items-center h-full">
          <img v-if="square21 == true" src="../assets/circlePlay.svg" alt="" />
          <img
            class="w-[48px]"
            v-else-if="square21 == false"
            src="../assets/crossPlay.svg"
            alt=""
          />
        </div>
      </div>
      <div
        class="cursor-pointer w-[80px] h-[80px] flex items-center justify-center border-black border-b-4 border-r-4"
        @click="getPosition([2, 2])"
      >
        <div class="flex m-2 justify-center items-center h-full">
          <img v-if="square22 == true" src="../assets/circlePlay.svg" alt="" />
          <img
            class="w-[48px]"
            v-else-if="square22 == false"
            src="../assets/crossPlay.svg"
            alt=""
          />
        </div>
      </div>
      <div
        class="cursor-pointer w-[80px] h-[80px] flex items-center justify-center border-black border-b-4"
        @click="getPosition([2, 3])"
      >
        <div class="flex m-2 justify-center items-center h-full">
          <img v-if="square23 == true" src="../assets/circlePlay.svg" alt="" />
          <img
            class="w-[48px]"
            v-else-if="square23 == false"
            src="../assets/crossPlay.svg"
            alt=""
          />
        </div>
      </div>
      <div
        class="cursor-pointer w-[80px] h-[80px] flex items-center justify-center border-black border-r-4"
        @click="getPosition([3, 1])"
      >
        <div class="flex m-2 justify-center items-center h-full">
          <img v-if="square31 == true" src="../assets/circlePlay.svg" alt="" />
          <img
            class="w-[48px]"
            v-else-if="square31 == false"
            src="../assets/crossPlay.svg"
            alt=""
          />
        </div>
      </div>
      <div
        class="cursor-pointer w-[80px] h-[80px] flex items-center justify-center border-black border-r-4"
        @click="getPosition([3, 2])"
      >
        <div class="flex m-2 justify-center items-center h-full">
          <img v-if="square32 == true" src="../assets/circlePlay.svg" alt="" />
          <img
            class="w-[48px]"
            v-else-if="square32 == false"
            src="../assets/crossPlay.svg"
            alt=""
          />
        </div>
      </div>
      <div
        class="cursor-pointer w-[80px] h-[80px] flex items-center justify-center border-black"
        @click="getPosition([3, 3])"
      >
        <div class="flex m-2 justify-center items-center h-full">
          <img v-if="square33 == true" src="../assets/circlePlay.svg" alt="" />
          <img
            class="w-[48px]"
            v-else-if="square33 == false"
            src="../assets/crossPlay.svg"
            alt=""
          />
        </div>
      </div>
    </div>

    <div
      v-if="playerBool"
      class="mt-6 border px-2 py-1 bg-blue-200 font-bold text-blue-500 rounded-2xl border-blue-400"
    >
      <div class="flex gap-1 justify-center items-center">
        <div>
          <img class="w-5" src="../assets/circle.svg" alt="circle" />
        </div>
        <div class="mr-1">
          {{ player }}
        </div>
      </div>
    </div>
    <div
      v-else
      class="mt-6 border px-2 py-1 bg-red-200 font-bold text-red-500 rounded-2xl border-red-400"
    >
      <div class="flex gap-1 justify-center items-center">
        <div>
          <img class="w-5" src="../assets/cross.svg" alt="circle" />
        </div>
        <div class="mr-1">
          {{ player }}
        </div>
      </div>
    </div>
  </div>
</template>

<style></style>
