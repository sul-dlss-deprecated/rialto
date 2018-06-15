<template>
  <el-container>
    <OrganizationList />
    <OrganizationTabularResults />
  </el-container>
</template>

<script>
import OrganizationList from './OrganizationList.vue'
import OrganizationTabularResults from './OrganizationTabularResults.vue'

export default {
  name: 'DepartmentsDashboard',
  components: {
    OrganizationList,
    OrganizationTabularResults
  },
  props: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
