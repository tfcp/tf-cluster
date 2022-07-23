<template>
  <div class="app-container">
    <el-form ref="configRule" :model="form" :rules="rules" label-width="120px">
      <el-form-item label="集群名称" prop="name">
        <el-input v-model="form.name" placeholder="请填写名称" />
      </el-form-item>
      <el-form-item label="集群ID" prop="cluster_id">
        <el-input v-model="form.cluster_id" placeholder="请填写名称" />
      </el-form-item>
      <el-form-item label="集群配置" prop="config">
        <el-input v-model="form.config" type="textarea" :rows="10" placeholder="请填写配置" />
      </el-form-item>
      <el-form-item label="集群介绍">
        <el-input v-model="form.introduction" type="textarea" :rows="10" placeholder="请填写集群描述"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
        <el-button @click="onClear">清除</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { saveConfig, getDetail } from '@/api/cluster'

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
      getDetail(this.$route.query.id).then(response => {
        const res = response.data.result
        this.form.id = res.id
        this.form.name = res.name
        this.form.cluster_id = res.cluster_id
        this.form.config = res.config
        this.form.create_at = res.create_at
        this.form.introduction = res.introduction
      })
    }
  },
  methods: {
    onSubmit() {
      this.$refs['configRule'].validate((valid) => {
        if (valid) {
          saveConfig(this.form).then(response => {
            this.$message({
              message: '提交成功',
              type: 'success'
            })
            this.$router.push('/cluster/config-list')
          })
        } else {
          this.$message({
            message: '参数有误!',
            type: 'warning'
          })
          return false
        }
      })
    },
    onClear() {
      this.form.name = ''
      this.form.cluster_id = ''
      this.form.introduction = ''
      this.$message({
        message: '清除成功!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

