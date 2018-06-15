<template>
  <div id="app">
    <nav>
      <a href="https://library.stanford.edu"><img src="./assets/sul-logo.svg" alt="Stanford Libraries" height="25"></a>
    </nav>
    <el-header>
      <h1>RIALTO Research at Stanford</h1>
    </el-header>
    <DepartmentsDashboard />
  </div>
</template>

<script>
import DepartmentsDashboard from './components/DepartmentsDashboard.vue'

export default {
  name: 'app',
  components: {
    DepartmentsDashboard
  }
}
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
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
