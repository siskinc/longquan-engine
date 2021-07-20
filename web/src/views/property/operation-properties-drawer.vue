<template>
  <el-drawer :title="title" :visible.sync="drawer" direction="rtl" size="80%">
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
</template>

<script>
export default {
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
    };
  },
  computed: {
    title: {
      get: function () {
        return `${this.selectData.code} 属性集`;
      },
    },
  },
};
</script>
