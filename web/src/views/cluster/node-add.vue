<template>
  <div class="app-container">
    <el-form :disabled="true" ref="configRule" :model="form" :rules="rules" label-width="120px">
      <el-form-item label="IP" prop="ip">
        <el-input v-model="form.ip" placeholder="请填写ip" />
      </el-form-item>
      <el-form-item label="集群名称" prop="cluster">
        <el-input v-model="form.cluster" placeholder="请填写集群名称" />
      </el-form-item>
      <el-form-item label="节点标签">
        <el-input v-model="form.labels" type="textarea" :rows="10" placeholder="请填写集群描述"/>
      </el-form-item>
      <el-form-item label="可否调度">
        <el-switch
          v-model="form.unschedulable"
          :active-value="1"
          :inactive-vlaue="2"
          active-color="#13ce66"
          inactive-color="#ff4949">
        </el-switch>
      </el-form-item>
      <el-form-item label="创建时间">
        <el-input v-model="form.create_at" type="date"/>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getNodeDetail } from '@/api/cluster'

export default {
  data() {
    return {
      form: {
        id: 0,
        name: '',
        labels: '',
        cluster: ''
      }
    }
  },
  created() {
    if (this.$route.query.id != null) {
      getNodeDetail(this.$route.query.id).then(response => {
        const res = response.data.result
        this.form.id = res.id
        this.form.ip = res.ip
        this.form.labels = res.labels
        this.form.cluster = res.cluster
        this.form.unschedulable = res.unschedulable
        this.form.create_at = res.create_at
      })
    }
  },
  methods: {
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

