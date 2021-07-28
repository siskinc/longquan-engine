<template>
  <div>
    <el-drawer
      v-loading="loadingPropertySetMap"
      :title="title"
      :visible.sync="drawer"
      direction="rtl"
      size="80%"
    >
      <el-button id="create-button" type="primary" @click="onCreate">
        创 建
      </el-button>
      <el-tree
        :data="propertiesTree"
        node-key="id"
        default-expand-all
        :expand-on-click-node="false"
      >
        <span slot-scope="{ node, data }" class="custom-tree-node">
          <span>{{ node.label }}</span>
          <span>{{ data }}</span>
        </span>
      </el-tree>
    </el-drawer>
    <create-property-form-dialog
      :visible.sync="dialogVisible.create"
      :namespace-id.sync="selectData.namespace_id"
      :property-set-code.sync="selectData.code"
      @refresh="refreshPropertySetMap"
    />
  </div>
</template>

<script>
import createPropertyFormDialog from "./create-property-dialog";

export default {
  components: {
    createPropertyFormDialog,
  },
  setup() {},
  props: {
    drawer: {
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
  data() {
    return {
      propertiesTree: [],
      defaultProps: {
        children: "children",
        label: "code",
      },
      dialogVisible: {
        create: false,
      },
      loadingPropertySetMap: false,
    };
  },
  computed: {
    title: {
      get: function () {
        return `${this.selectData.code} 属性集`;
      },
    },
  },
  watch: {
    drawer: async function (newValue, _) {
      if (newValue) {
        await this.refreshPropertySetMap();
      }
    },
  },
  async created() {
    await this.refreshPropertySetMap();
  },
  methods: {
    refreshPropertySetMap: async function () {
      this.loadingPropertySetMap = true;
      this.propertiesTree = [];
      this.loadingPropertySetMap = false;
    },
    onCreate: function () {
      this.dialogVisible.create = true;
    },
  },
};
</script>

<style scoped>
#create-button {
  margin-left: 5px;
}
</style>
