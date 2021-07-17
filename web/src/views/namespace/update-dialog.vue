<template>
  <el-dialog title="更新命名空间" :visible.sync="dialogVisible">
    <el-form :model="form">
      <el-form-item label="ID" :label-width="formLabelWidth">
        <el-input
          v-model="selectData.id"
          autocomplete="off"
          disabled
          class="display-data"
        />
      </el-form-item>
      <el-form-item label="名称" :label-width="formLabelWidth">
        <el-input
          v-model="selectData.name"
          autocomplete="off"
          disabled
          class="display-data"
        />
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
      <el-button type="primary" @click="onSubmit"> 更 新 </el-button>
    </div>
  </el-dialog>
</template>

<script>
import { updateNamespace } from "@/api/namespace";

export default {
  setup() {},
  props: {
    visible: {
      type: Boolean,
      default: false,
    },
    selectData: {
      type: Object,
      default: {},
    },
  },
  data: function () {
    return {
      formLabelWidth: "80px",
      form: {
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
  watch: {
    selectData(newVal, _) {
      this.form.description = newVal.description;
    },
  },
  methods: {
    onSubmit: function () {
      updateNamespace(this.selectData.id, this.form).then((_) => {
        this.$message.success(`更新 ${this.selectData.name} 命名空间成功`);
        this.dialogVisible = false;
        this.$emit("refresh");
      });
    },
  },
};
</script>
