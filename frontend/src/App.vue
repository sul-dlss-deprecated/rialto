<template>
  <div id="app">
    <nav>
      <a href="https://library.stanford.edu"><img src="./assets/sul-logo.svg" alt="Stanford Libraries" height="25"></a>
    </nav>
    <el-header>
      <h1>RIALTO Research at Stanford</h1>
    </el-header>
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'app',
  components: {
    HelloWorld
  }
}
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.el-header {
  background-color: #8c1515;
  color: white;
  h1 {
    line-height: 60px;
    margin: 0;
    padding: 0;
  }
}

</style>
