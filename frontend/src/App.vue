<script setup>
  import HeaderVue from "./components/HeaderVue.vue"
  import HashGame from "./components/HashGame.vue"
</script>

<template>
  <head>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=DM+Sans:ital,opsz,wght@0,9..40,100..1000;1,9..40,100..1000&display=swap" rel="stylesheet">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=DM+Sans:ital,opsz,wght@0,9..40,100..1000;1,9..40,100..1000&family=Lobster&family=Lobster+Two:ital,wght@0,400;0,700;1,400;1,700&display=swap" rel="stylesheet">
  </head>
  <body>
    <div id="body" class="flex flex-col gap-3">
      <HeaderVue></HeaderVue>
      <HashGame></HashGame>
    </div>
  </body>
</template>

<style>
  * {
  box-sizing: border-box;
  padding: 0;
  margin: 0;
  overflow: hidden;
  font-family: "DM Sans", sans-serif;
  }
</style>