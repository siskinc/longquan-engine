<template>
  <el-dialog title="创建新的属性" :visible.sync="dialogVisible">
    <el-form :model="form" :label-width="formLabelWidth">
      <!-- 命名空间 -->
      <el-form-item label="命名空间">
        <el-input v-model="namespaceName" autocomplete="off" disabled />
      </el-form-item>

      <!-- 属性集 -->
      <el-form-item label="属性集">
        <el-input v-model="propertySetCode" autocomplete="off" disabled />
      </el-form-item>

      <!-- 填写名称 -->
      <el-form-item label="名称【英文】">
        <el-input v-model="form.name" autocomplete="off" />
      </el-form-item>

      <!-- 字段分类 -->
      <el-form-item label="字段分类">
        <el-select v-model="form.class" placeholder="请选择字段分类">
          <el-option
            v-for="propertyTypeEnum in propertyTypeEnumList"
            :key="propertyTypeEnum.code"
            :label="propertyTypeEnum.name"
            :value="propertyTypeEnum.code"
          />
        </el-select>
      </el-form-item>

      <!-- 字段类型 -->
      <el-form-item label="字段类型">
        <el-select v-model="form.type" placeholder="请选择字段类型">
          <el-option
            v-for="classType in classTypeList"
            :key="classType"
            :label="classType"
            :value="classType"
          />
        </el-select>
      </el-form-item>

      <!-- 描述 -->
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
import { queryPropertySet } from "@/api/property-set";
import {
  createProperty,
  queryPropertyTypeEnum,
  queryPropertySystemClassEnum,
} from "@/api/property";
import _ from "lodash";

export default {
  props: {
    visible: {
      type: Boolean,
      default: false,
    },
    namespaceId: {
      type: String,
      default: "",
    },
    propertySetCode: {
      type: String,
      default: "",
    },
  },
  data: function () {
    return {
      formLabelWidth: "120px",
      form: {
        name: "",
        class: "",
        type: "",
        description: "",
      },
      loadingNamespace: true,
      loadingPropertySet: true,
      namespaceName: "",
      propertySetName: "",
      namespaceList: [],
      propertySetList: [],
      propertyTypeEnumList: [],
      classTypeList: [],
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
  watch: {
    namespaceId: async function () {
      await this.queryNamespaceDataList();
    },
    propertySetCode: async function () {
      await this.queryPropertySetDataList();
    },
    "form.class": async function () {
      await this.queryPropertyTypeList();
    },
  },
  async created() {
    await this.queryPropertyTypeEnumList();
    await this.queryNamespaceDataList();
  },
  methods: {
    queryNamespaceDataList: async function () {
      this.loadingNamespace = true;
      this.namespaceList = (await queryNamespaceDataList()) || [];
      for (let i = 0; i < this.namespaceList.length; i++) {
        const item = this.namespaceList[i];
        if (item.id === this.namespaceId) {
          this.namespaceName = item.name;
          break;
        }
      }
      this.loadingNamespace = false;
    },
    queryPropertySetDataList: async function () {
      this.loadingPropertySet = true;
      this.propertySetList =
        (await queryPropertySet({ namespace_id: this.namespaceId })) || [];
      this.loadingPropertySet = false;
    },
    queryPropertyTypeEnumList: async function () {
      const data = (await queryPropertyTypeEnum()) || {};
      this.propertyTypeEnumList = [];
      for (const key in data) {
        if (Object.hasOwnProperty.call(data, key)) {
          const element = data[key];
          this.propertyTypeEnumList.push({ code: key, name: element });
        }
      }
    },
    queryPropertyTypeList: async function () {
      this.classTypeList = [];
      switch (this.form.class) {
        case "system":
          this.classTypeList = (await queryPropertySystemClassEnum()) || [];
          break;
        case "custom":
          for (let i = 0; i < this.propertySetList.length; i++) {
            const element = this.propertySetList[i];
            this.classTypeList.push(element.code);
          }
          break;
        default:
          this.$message.error("未知字段分类");
          break;
      }
      console.log(`this.classTypeList: ${JSON.stringify(this.classTypeList)}`);
    },
    onSubmit: function () {
      let data = _.cloneDeep(this.form);
      data.namespace_id = this.namespaceId;
      data.property_set_code = this.propertySetCode;
      console.log(data);
      createProperty(data).then((_) => {
        this.$message.success(`创建 ${this.form.name} 属性成功`);
        this.dialogVisible = false;
        this.$emit("refresh");
      });
    },
  },
};
</script>
