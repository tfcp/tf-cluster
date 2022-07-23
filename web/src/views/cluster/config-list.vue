<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true">
        <el-row>
          <el-form-item label="集群名称">
            <el-input v-model.trim="searchParams.name"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="el-icon-search" @click="fetchData()">搜索</el-button>
            <el-button type="success" icon="el-icon-s-platform" @click="toEdit(0)">创建新集群</el-button>
          </el-form-item>
        </el-row>
      </el-form>
    </div>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column
        type="index"
        align="center"
        label="ID"
        width="50">
      </el-table-column>
      <el-table-column label="集群名称" align="center">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="集群ID" align="center">
        <template slot-scope="scope">
          {{ scope.row.cluster_id }}
        </template>
      </el-table-column>
      <el-table-column label="介绍"  align="center">
        <template slot-scope="scope">
          {{ scope.row.introduction }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="create_at" label="创建日期" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.create_at | formatTime}}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作" width="220">
        <template slot-scope="scope">
          <el-row>
            <el-button type="info" icon="el-icon-edit" @click="toEdit(scope.row.id)">编辑</el-button>
          </el-row>
        </template>
      </el-table-column>
    </el-table>
    <div class="block">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :page-sizes="[1, 10, 50]"
        layout="total, sizes, prev, pager, next, jumper"
        :total="countData">
      </el-pagination>
    </div>
  </div>
</template>

<script>
import { getConfigList, getConfigCount } from '@/api/cluster'

export default {
  filters: {
  },
  data() {
    return {
      searchParams: {
        page_size: 20,
        page: 1,
        name: '',
      },
      list: null,
      countData: 0,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getConfigList(this.searchParams).then(response => {
        this.list = response.data.result
        this.listLoading = false
      })
      getConfigCount(this.searchParams).then(response => {
        this.countData = response.data.result
      })
    },
    handleSizeChange(val) {
      this.searchParams.page_size = val
      this.fetchData()
    },
    handleCurrentChange(val) {
      this.searchParams.page = val
      this.fetchData()
    },
    refresh() {
      this.fetchData(() => {
        this.$message.success('刷新成功')
      })
    },
    toEdit(id) {
      let path = ''
      if (id === 0) {
        path = '/cluster/config-add'
      } else {
        path = '/cluster/config-add?id=' + id
      }
      this.$router.push(path)
    }
  }
}
</script>
