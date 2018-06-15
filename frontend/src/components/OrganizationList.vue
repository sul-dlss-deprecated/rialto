<template>
  <el-aside>
    <el-menu default-active="1" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose">
      <el-submenu index="1">
        <template slot="title">
          <span>Business</span>
        </template>
        <el-menu-item index="1-1">
          <span>Marketing</span>
        </el-menu-item>
      </el-submenu>

      <el-submenu index="2">
        <template slot="title">
          <span>Education</span>
        </template>
        <el-menu-item index="1-1">
          <span>Elementary Education</span>
        </el-menu-item>
      </el-submenu>

      <el-submenu index="3">
        <template slot="title">
          <span>Engineering</span>
        </template>
        <el-menu-item index="1-1">
          <span>Civil Engineering</span>
        </el-menu-item>
      </el-submenu>

      <el-submenu index="4">
        <template slot="title">
          <span>Humanities &amp; Science</span>
        </template>
        <el-menu-item index="1-1">
          <span>Anthropology</span>
        </el-menu-item>
        <el-menu-item index="1-1">
          <span>Applied Physics</span>
        </el-menu-item>
        <el-menu-item index="1-1">
          <span>Art &amp; Art History</span>
        </el-menu-item>
        <el-menu-item index="1-1">
          <span>Biology</span>
        </el-menu-item>
      </el-submenu>
    </el-menu>
  </el-aside>
</template>

<script>
export default {
  name: 'OrganizationList',
}
</script>

<style lang="scss">
.el-submenu.is-opened > .el-submenu__title {
   background-color: #e1ded4;
}

.el-submenu__title:hover, .el-menu-item:hover {
   background-color: #e1dfd6;
}
</style>
