<template>
  <v-dialog v-model="dialog" max-width="550px">
    <template v-slot:activator="{ on }">
      <v-btn color="orange white--text" v-on="on" large fixed bottom right fab>
        <v-icon>mdi-cloud-upload</v-icon>
      </v-btn>
    </template>

    <v-card>
      <v-card-title>上传APK文件</v-card-title>
      <v-card-text>
        <v-container>
          <v-form ref="form" v-model="valid" lazy-validation>
            <v-text-field
              prepend-icon="mdi-android"
              v-model="name"
              :counter="10"
              :rules="nameRules"
              label="请输入上传者名称"
              hint="请输入上传者名称"
              clearable
              required
            ></v-text-field>
            <v-textarea
              v-model="description"
              outlined
              prepend-icon="mdi-pen"
              :counter="50"
              :rules="descriptionRules"
              clearable
              label="请输入描述信息"
              hint="请输入描述信息"
            ></v-textarea>
            <v-file-input
              v-model="apk"
              chips
              prepend-icon="mdi-file"
              accept=".apk"
              label="选择文件"
              hint="请选择APK文件"
              show-size
              :rules="apkRules"
              required
            ></v-file-input>
          </v-form>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>

        <v-btn color="grey darken-1" text @click="dialog = false">
          取消
        </v-btn>

        <v-btn :disabled="!valid" color="primary" text @click="submit">
          确定
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { mapActions } from "vuex";

export default {
  data() {
    return {
      dialog: false,
      valid: true,
      name: "",
      nameRules: [
        (v) => !!v || "上传者不可为空",
        (v) => (v && v.length <= 10) || "上传者名称不得超过10个字符",
      ],
      description: "",
      descriptionRules: [(v) => v.length <= 50 || "描述信息不得超过50个字符"],
      apk: null,
      apkRules: [(v) => !!v || "apk不可为空"],
    };
  },
  methods: {
    ...mapActions(["addFile"]),
    submit() {
      this.dialog = false;
      this.addFile({
        file: this.apk,
        name: this.name,
        description: this.description,
      });
    },
  },
};
</script>
