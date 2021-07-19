<template>
  <el-dialog title="创建新的属性集" :visible.sync="dialogVisible">
    <el-form :model="form" :label-width="formLabelWidth">
      <el-form-item label="命名空间">
        <el-select
          v-model="form.namespace_id"
          v-loading="loadingNamespace"
          placeholder="请选择命名空间"
        >
          <el-option
            v-for="namespace in namespaceList"
            :key="namespace.id"
            :label="namespace.name"
            :value="namespace.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="名称【英文】">
        <el-input v-model="form.code" autocomplete="off" />
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
      <el-button type="primary" @click="onSubmit"> 创 建 </el-button>
    </div>
  </el-dialog>
</template>

<script>
import { queryNamespaceDataList } from "@/api/namespace";
import { createPropertySet } from "@/api/property-set";

export default {
  props: {
    visible: {
      type: Boolean,
      default: false,
    },
  },
  data: function () {
    return {
      formLabelWidth: "120px",
      form: {
        code: "",
        description: "",
        namespace_id: "",
      },
      loadingNamespace: true,
      namespaceList: [],
    };
  },
  computed: {
    dialogVisible: {
      get() {
        return this.visible;
      },
      async set(val) {
        this.$emit("update:visible", val);
        await this.queryNamespaceDataList();
      },
    },
  },
  async created() {
    await this.queryNamespaceDataList();
  },
  methods: {
    queryNamespaceDataList: async function () {
      this.loadingNamespace = true;
      this.namespaceList = await queryNamespaceDataList();
      this.loadingNamespace = false;
    },
    onSubmit: function () {
      console.log(this.form);
      createPropertySet(this.form).then((_) => {
        this.$message.success(`创建 ${this.form.code} 属性集成功`);
        this.dialogVisible = false;
        this.$emit("refresh");
      });
    },
  },
};
</script>
