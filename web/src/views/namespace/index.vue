<template>
  <div class="app-container">
    <!-- 表单处，查询命名空间 -->
    <el-form ref="form" :model="form" label-width="80px" :inline="true">
      <el-button
        type="primary"
        icon="el-icon-circle-plus-outline"
        @click="dialogVisible.create = true"
      />
      <el-form-item label="ID">
        <el-input v-model="form.id" placeholder="ID" clearable />
      </el-form-item>
      <el-form-item label="名称">
        <el-input v-model="form.name" placeholder="名称" clearable />
      </el-form-item>
      <el-form-item label="描述">
        <el-input v-model="form.description" placeholder="描述" clearable />
      </el-form-item>
      <!-- 批量操作 -->
      <el-button-group>
        <el-button
          type="primary"
          icon="el-icon-search"
          @click="refreshNamespaceDataList"
        />
        <el-button
          type="danger"
          icon="el-icon-delete"
          @click="handleBatchDeleteNamespace"
        />
      </el-button-group>
    </el-form>
    <!-- 表格  展示命名空间 -->
    <el-table
      v-loading="loading"
      :data="namespaceDataList"
      element-loading-text="Loading..."
      border
      fit
      highlight-current-row
      @selection-change="handleSelectionChange"
      @cell-dblclick="handleUpdateNamespace"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column align="center" label="ID" width="240" fixed="left">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="名称" width="250">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="描述" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.description }}</span>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页 -->
    <div class="block" align="center">
      <el-pagination
        :current-page="pageInfo.pageIndex"
        :page-sizes="[10, 20, 30, 40, 50, 60, 70, 80, 90, 100]"
        :page-size="pageInfo.pageSize"
        :total="namespaceTotal"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
    <!-- 创建命名空间 -->
    <create-namespace-form-dialog
      :visible.sync="dialogVisible.create"
      @refresh="refreshNamespaceDataList"
    />
    <!-- 更新命名空间 -->
    <update-namespace-form-dialog
      :visible.sync="dialogVisible.update"
      :select-data.sync="selectedData"
      @refresh="refreshNamespaceDataList"
    />
  </div>
</template>

<script>
import { queryNamespace, deleteNamespace } from "@/api/namespace";
import createNamespaceFormDialog from "./create-dialog";
import updateNamespaceFormDialog from "./update-dialog";
export default {
  components: {
    createNamespaceFormDialog,
    updateNamespaceFormDialog,
  },
  data() {
    return {
      form: {
        id: "",
        name: "",
        description: "",
      },
      pageInfo: {
        pageIndex: 1,
        pageSize: 10,
      },
      loading: false,
      namespaceDataList: [],
      namespaceTotal: 0,
      selectedRows: [],
      dialogVisible: {
        create: false,
        update: false,
      },
      selectedData: {},
    };
  },
  created: function () {
    this.refreshNamespaceDataList();
  },
  methods: {
    refreshNamespaceDataList: function () {
      this.loading = true;
      const params = {
        id: this.form.id,
        name: this.form.name,
        description: this.form.description,
        page_index: this.pageInfo.pageIndex,
        page_size: this.pageInfo.pageSize,
      };
      queryNamespace(params).then((response) => {
        this.namespaceDataList = response.data;
        this.namespaceTotal = response.total;
        this.loading = false;
      });
    },
    handleSelectionChange: function (selectedRows) {
      this.selectedRows = selectedRows;
    },
    handleBatchDeleteNamespace: async function () {
      for (const i in this.selectedRows) {
        const _id = this.selectedRows[i].id;
        const name = this.selectedRows[i].name;
        await deleteNamespace(_id).then((_) => {
          this.$message.success(`删除 "${name}" 命名空间成功`);
        });
      }
      this.refreshNamespaceDataList();
    },
    handleCurrentChange: function (currentIndex) {
      this.pageInfo.pageIndex = currentIndex;
      this.refreshNamespaceDataList();
    },
    handleSizeChange: function (currentSize) {
      this.pageInfo.pageSize = currentSize;
      this.refreshNamespaceDataList();
    },
    handleUpdateNamespace(row) {
      this.selectedData = row;
      this.dialogVisible.update = true;
    },
  },
};
</script>
