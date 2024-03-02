<script setup>
    import { ref } from 'vue';

    const player = ref("Player 1")
    function getPosition(position){
        fetch("http://localhost:8081/v1/hash/position", {
            method: 'POST',
            body: JSON.stringify({
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
        {{ player }}
    </div>
      
</template>

<style>

</style>