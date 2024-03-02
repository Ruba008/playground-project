<script setup>
    import { ref } from 'vue';
    const player = ref("Player 1")
    // Defining if is the player's turn 1 or 2
    // true: player 1
    // false: player 2
    const playerBool = ref(true)

    function getPosition(position){
        playerBool.value = playerBool.value ? false : true
        player.value = playerBool.value ? "Player 1" : "Player 2"
        fetch("http://localhost:8081/v1/hash/position", {
            method: 'POST',
            body: JSON.stringify({
                player: playerBool.value,
                position: position
            })
        })
        .then(res => console.log(res.json()))
    }
</script>

<template>
    <div class="flex w-full justify-center flex-col items-center">
        <div class="grid grid-cols-3 text-center w-60 h-60 mt-10">
            <div class="flex cursor-pointer items-center justify-center border-r-4 border-b-4 border-black" @click="getPosition([1,1])"></div>
            <div class="cursor-pointer flex items-center justify-center border-r-4 border-b-4 border-black" @click="getPosition([1,2])"></div>
            <div class="cursor-pointer flex items-center justify-center border-black border-b-4" @click="getPosition([1,3])"></div>
            <div class="cursor-pointer flex items-center justify-center border-black border-b-4 border-r-4" @click="getPosition([2,1])"></div>
            <div class="cursor-pointer flex items-center justify-center border-black border-b-4 border-r-4" @click="getPosition([2,2])"></div>
            <div class="cursor-pointer flex items-center justify-center border-black border-b-4" @click="getPosition([2,3])"></div>
            <div class="cursor-pointer flex items-center justify-center border-black border-r-4" @click="getPosition([3,1])"></div>
            <div class="cursor-pointer flex items-center justify-center border-black border-r-4" @click="getPosition([3,2])"></div>
            <div class="cursor-pointer flex items-center justify-center border-black" @click="getPosition([3,3])"></div>
        </div>
        <div v-if="playerBool" class="mt-6 border px-2 py-1 bg-blue-200 font-bold text-blue-500 rounded-2xl border-blue-400">
            <div class="flex gap-1 justify-center items-center">
                <div>
                    <img class="w-5" src="../assets/circle.svg" alt="circle">
                </div>
                <div class="mr-1">
                    {{ player }}
                </div>
            </div>
        </div>
        <div v-else class="mt-6 border px-2 py-1 bg-red-200 font-bold text-red-500 rounded-2xl border-red-400">
            <div class="flex gap-1 justify-center items-center">
                <div>
                    <img class="w-5" src="../assets/cross.svg" alt="circle">
                </div>
                <div class="mr-1">
                    {{ player }}
                </div>
            </div>
        </div>
    </div>
      
</template>

<style>

</style>