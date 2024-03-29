<template>
  <div class="app-container">
    <!-- 表单，查询属性集 -->
    <el-form ref="form" :model="form" label-width="120px" :inline="true">
      <el-button
        type="primary"
        icon="el-icon-circle-plus-outline"
        @click="dialogVisible.create = true"
      />
      <el-form-item label="属性集ID">
        <el-input v-model="form.id" placeholder="ID" clearable />
      </el-form-item>
      <el-form-item label="命名空间">
        <el-select
          v-model="form.namespace_id"
          v-loading="loadingNamespace"
          placeholder="请选择命名空间"
        >
          <el-option
            key=""
            label="查询所有命名空间"
            value=""
            style="color: red"
          />
          <el-option
            v-for="namespace in namespaceList"
            :key="namespace.id"
            :label="namespace.name"
            :value="namespace.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="属性集名称">
        <el-input v-model="form.code" placeholder="属性集名称" clearable />
      </el-form-item>
      <el-form-item label="属性集描述">
        <el-input
          v-model="form.description"
          placeholder="属性集描述"
          clearable
        />
      </el-form-item>
      <!-- 批量操作 -->
      <el-button-group>
        <el-button
          type="primary"
          icon="el-icon-search"
          @click="refreshPropertySetDataList"
        />
        <el-button
          type="danger"
          icon="el-icon-delete"
          @click="handleBatchDeletePropertySet"
        />
      </el-button-group>
    </el-form>

    <!-- 表格展示 -->
    <el-table
      v-loading="loadingPropertySetDataList"
      :data="propertySetList"
      element-loading-text="Loading..."
      border
      fit
      highlight-current-row
      @selection-change="handleSelectionChange"
      @cell-dblclick="handleUpdatePropertySet"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column align="center" label="ID" width="240" fixed="left">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="命名空间" width="250">
        <template slot-scope="scope">
          {{ scope.row.namespace_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="名称" width="250">
        <template slot-scope="scope">
          {{ scope.row.code }}
        </template>
      </el-table-column>
      <el-table-column label="描述" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.description }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <el-button type="text" @click="handleUpdatePropertySet(scope.row)">修改</el-button>
          <el-button type="text" @click="operationProperties(scope.row)">属性字段</el-button>
        </template>
      </el-table-column>
    </el-table>
    <create-property-set-form-dialog
      :visible.sync="dialogVisible.create"
      @refresh="refreshPropertySetDataList"
    />
    <update-property-set-form-dialog
      :visible.sync="dialogVisible.update"
      :select-data.sync="selectData"
      @refresh="refreshPropertySetDataList"
    />
    <operation-properties-drawer
      :drawer.sync="drawer"
      :select-data.sync="selectData"
    />
  </div>
</template>

<script>
import { queryNamespaceDataList } from "@/api/namespace";
import { queryPropertySet, deletePropertySet } from "@/api/property-set";
import createPropertySetFormDialog from "./create-property-set-dialog";
import updatePropertySetFormDialog from "./update-property-set-dialog";
import operationPropertiesDrawer from "./operation-properties-drawer"

export default {
  components: {
    createPropertySetFormDialog,
    updatePropertySetFormDialog,
    operationPropertiesDrawer,
  },
  setup() {},
  data() {
    return {
      form: {
        id: "",
        namespace_id: "",
        code: "",
        description: "",
      },
      dialogVisible: {
        create: false,
        update: false,
      },
      propertySetList: [],
      loadingPropertySetDataList: true,
      namespaceList: {},
      loadingNamespace: true,
      selectedDataList: [],
      selectData: {},
      drawer: false,
    };
  },
  async created() {
    await this.queryNamespaceList();
    await this.refreshPropertySetDataList();
  },
  methods: {
    queryNamespaceList: async function () {
      this.loadingNamespace = true;
      this.namespaceList = await queryNamespaceDataList();
      this.loadingNamespace = false;
    },
    refreshPropertySetDataList: async function () {
      this.loadingPropertySetDataList = true;

      // 构建命名空间的id和信息映射
      const namespaceMap = {};

      for (const index in this.namespaceList) {
        const namespace = this.namespaceList[index];
        namespaceMap[namespace.id] = namespace;
      }

      // 获取属性集
      this.propertySetList = await queryPropertySet(this.form);

      // 补全属性集的基本信息
      for (const index in this.propertySetList) {
        const propertySet = this.propertySetList[index];
        const namespace_id = propertySet.namespace_id;
        propertySet.namespace_name = namespaceMap[namespace_id].name;
      }

      this.loadingPropertySetDataList = false;
    },
    handleBatchDeletePropertySet: async function () {
      for (const index in this.selectedDataList) {
        const propertySet = this.selectedDataList[index];
        await deletePropertySet(propertySet.id).then((_) => {
          this.$message.success(`删除 ${propertySet.code} 属性成功`);
        });
      }
      this.refreshPropertySetDataList();
    },
    handleSelectionChange: function (selectedDataList) {
      this.selectedDataList = selectedDataList;
    },
    handleUpdatePropertySet: function (row) {
      this.selectData = row;
      this.dialogVisible.update = true;
    },
    operationProperties: function (row) {
      this.selectData = row;
      this.drawer = true;
    },
  },
};
</script>
