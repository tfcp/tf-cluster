<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true">
        <el-row>
          <el-form-item label="IP名称">
            <el-input v-model.trim="searchParams.ip"></el-input>
          </el-form-item>
          <el-form-item label="集群列表">
            <el-select v-model.trim="searchParams.cluster">
              <el-option
                v-for="item in clusterList"
                :key="item.value"
                :label="item.label"
                :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="el-icon-search" @click="fetchData()">搜索</el-button>
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
      <el-table-column label="节点名称" align="center">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="关闭调度" width="110" align="center">
        <template slot-scope="scope">
          <el-tag v-if= "scope.row.unschedulable == 1">否</el-tag>
          <el-tag v-else type="warning">是</el-tag>
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
            <el-button type="info" icon="el-icon-edit-outline" @click="toEdit(scope.row.id)">查看</el-button>
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
import { getNodeList, getNodeCount } from '@/api/cluster'

export default {
  filters: {
    statusFilter(status) {
      var res = 'success'
      if (status == "2"){
        res = 'danger'
      }
      // deleted: 'gray'
      return res
    }
  },
  data() {
    return {
      searchParams: {
        page_size: 20,
        page: 1,
        name: '',
        role: '0'
      },
      clusterList: [
        {
          value: '0',
          label: '全部'
        }
      ],
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
      getNodeList(this.searchParams).then(response => {
        this.list = response.data.result
        this.listLoading = false
      })
      getNodeCount(this.searchParams).then(response => {
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
        path = '/cluster/node-add'
      } else {
        path = '/cluster/node-add?id=' + id
      }
      this.$router.push(path)
    }
  }
}
</script>
