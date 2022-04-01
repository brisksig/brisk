<template>
  <div id="app">
    <img src="./assets/logo.png">
    <router-view/>
  </div>
</template>

<script src="https://unpkg.com/axios/dist/axios.min.js"></script>

<script>
export default {
  name: 'App',
  mounted () {
    this.$axios
      .post('app01/api/')
      .then(response => (this.info = response))
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
