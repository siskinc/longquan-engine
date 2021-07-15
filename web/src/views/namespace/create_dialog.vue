<template>
  <el-dialog title="创建新的命名空间" :visible.sync="dialogVisible">
    <el-form :model="form">
      <el-form-item label="名称" :label-width="formLabelWidth">
        <el-input v-model="form.name" autocomplete="off" />
      </el-form-item>
      <el-form-item label="描述" :label-width="formLabelWidth">
        <el-input
          v-model="form.description"
          type="textarea"
          autocomplete="off"
        />
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogVisible = false">取 消</el-button>
      <el-button type="primary" @click="onSubmit"> 确 定 </el-button>
    </div>
  </el-dialog>
</template>

<script>
import { createNamespace } from "@/api/namespace";
export default {
  props: {
    visible: {
      type: Boolean,
      default: false,
    },
  },
  data: function () {
    return {
      formLabelWidth: "80px",
      form: {
        name: "",
        description: "",
      },
    };
  },
  computed: {
    dialogVisible: {
      get() {
        return this.visible;
      },
      set(val) {
        this.$emit("update:visible", val);
      },
    },
  },
  methods: {
    onSubmit: function () {
      createNamespace(this.form).then((_) => {
        this.$message.success(`创建 ${this.form.name} 命名空间成功`);
        this.dialogVisible = false;
        this.$emit("refresh");
      });
    },
  },
};
</script>
