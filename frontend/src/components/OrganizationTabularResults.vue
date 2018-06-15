<template>
  <el-main>
    <h2>{{ title }}</h2>
    <el-table :data="tableData3" height="250" style="width: 100%">
      <el-table-column prop="organization" label="Departments" width="180">
      </el-table-column>
      <el-table-column prop="faculty" label="Faculty" width="180">
      </el-table-column>
      <el-table-column prop="articles" label="Articles">
      </el-table-column>
      <el-table-column prop="citations" label="Citations">
      </el-table-column>
      <el-table-column prop="journals" label="Journals">
      </el-table-column>
      <el-table-column prop="grants" label="Grants">
      </el-table-column>
      <el-table-column prop="citations" label="Citations">
      </el-table-column>
    </el-table>
  </el-main>
</template>

<script>
export default {
  name: 'OrganizationTabularResults',
  data() {
    return {
      title: "Humanities & Sciences",
      tableData3: [{
          organization: 'Anthropology',
          faculty: '31',
          articles: '19',
          citations: '392',
          journals: '3',
          grants: '12',
          collaborations: '12'
        },
        {
          organization: 'Applied Physics',
          faculty: '31',
          articles: '19',
          citations: '392',
          journals: '3',
          grants: '12',
          collaborations: '12'
        },
        {
          organization: 'Art & Art History',
          faculty: '31',
          articles: '19',
          citations: '392',
          journals: '3',
          grants: '12',
          collaborations: '12'
        },
        {
          organization: 'Biology',
          faculty: '31',
          articles: '19',
          citations: '392',
          journals: '3',
          grants: '12',
          collaborations: '12'
        },
        {
          organization: 'Chemestry',
          faculty: '31',
          articles: '19',
          citations: '392',
          journals: '3',
          grants: '12',
          collaborations: '12'
        }
        ]
    }
  }
}
</script>

<style lang="scss">
</style>
