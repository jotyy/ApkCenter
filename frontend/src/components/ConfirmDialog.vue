<template>
  <v-dialog v-model="dialog" max-width="400">
    <template v-slot:activator="{ on }">
      <v-btn v-on="on">
        <v-icon color="error">mdi-close-circle</v-icon>
      </v-btn>
    </template>

    <v-card>
        <v-card-title class="headline4">你确定要删除这个文件吗?</v-card-title>

        <v-card-text>
          删除后文件将无法恢复，请确认后再操作。
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn
            color="gray darken-1"
            text
            @click="dialog = false"
          >
            否
          </v-btn>

          <v-btn
            color="error"
            text
            @click="deleteFile"
          >
            是
          </v-btn>
        </v-card-actions>
      </v-card>
  </v-dialog>
</template>

<script>
export default {
  props:["fileId"],
  data() {
    return {
      dialog: false
    }
  },
  methods: {
    deleteFile() {
      this.dialog = false;
      this.$emit("onDeleteFile", this.fileId)
    }
  }
}
</script>