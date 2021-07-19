<template>
  <el-dialog title="更新属性集" :visible.sync="dialogVisible">
    <el-form :model="form" :label-width="formLabelWidth">
      <el-form-item label="ID">
        <el-input
          v-model="selectData.id"
          autocomplete="off"
          disabled
          class="display-data"
        />
      </el-form-item>
      <el-form-item label="命名空间">
        <el-input
          v-model="selectData.namespace_name"
          autocomplete="off"
          disabled
        />
      </el-form-item>
      <el-form-item label="名称【英文】">
        <el-input
          v-model="selectData.code"
          autocomplete="off"
          disabled
          class="display-data"
        />
      </el-form-item>
      <el-form-item label="描述">
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
import { updatePropertySet } from "@/api/property-set";

export default {
  props: {
    visible: {
      type: Boolean,
      default: false,
    },
    selectData: {
      type: Object,
      default: () => {
        return {};
      },
    },
  },
  data: function () {
    return {
      formLabelWidth: "120px",
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
      updatePropertySet(this.selectData.id, this.form).then((_) => {
        this.$message.success(`更新 ${this.selectData.code} 属性集成功`);
        this.dialogVisible = false;
        this.$emit("refresh");
      });
    },
  },
};
</script>
