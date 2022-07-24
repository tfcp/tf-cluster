<template>
  <div class="app-container">
    <el-form ref="configRule" :model="form" :rules="rules" label-width="120px">
      <el-form-item label="IP" prop="ip">
        <el-input v-model="form.ip" placeholder="请填写ip" />
      </el-form-item>
      <el-form-item label="集群名称" prop="cluster_id">
        <el-input v-model="form.cluster_id" placeholder="请填写集群名称" />
      </el-form-item>
      <el-form-item label="节点标签">
        <el-input v-model="form.introduction" type="textarea" :rows="10" placeholder="请填写集群描述"/>
      </el-form-item>
      <el-form-item label="关闭调度">
        <el-input v-model="form.introduction" type="textarea" :rows="10" placeholder="是否关闭调度"/>
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
        cluster_id: '',
        config: '',
        introduction: ''
      },
      rules: {
        name: [
          { required: true, message: '请输入集群名称', trigger: 'blur' },
          { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' }
        ]
        // pwd: [
        //   { required: true, message: '请输入用户密码', trigger: 'blur' },
        //   { min: 6, max: 256, message: '长度在 6 到 128 个字符', trigger: 'blur' }
        // ]
        // age: [
        //   { required: true, message: '请填写年龄', trigger: 'blur' },
        //   { type: 'number', message: '请填写数字', trigger: 'blur', transform: (value) => Number(value)},
        // ]
      }
    }
  },
  created() {
    if (this.$route.query.id != null) {
      getNodeDetail(this.$route.query.id).then(response => {
        const res = response.data.result
        this.form.id = res.id
        this.form.ip = res.ip
        this.form.cluster = res.cluster
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

