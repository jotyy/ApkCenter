<template>
  <v-col id="download">
    <v-subheader>下载中心</v-subheader>
    <div v-if="sortFiles">
      <v-card flat v-for="project in sortFiles" :key="project.title">
        <v-layout row wrap :class="`pa-3 project complete`">
          <v-flex xs12 md2>
            <div class="caption grey--text">文件名</div>
            <div class="font-weight-bold">{{ project.file_name }}</div>
          </v-flex>
          <v-flex xs6 sm6 md4>
            <div class="caption grey--text">描述</div>
            <div>{{ project.apk_description }}</div>
          </v-flex>
          <v-flex xs6 sm6 md1>
            <div class="caption grey--text">上传者</div>
            <div>{{ project.upload_by }}</div>
          </v-flex>
          <v-flex xs6 sm6 md2>
            <div class="caption grey--text">上传时间</div>
            <div>{{ project.upload_time }}</div>
          </v-flex>
          <v-flex xs6 sm6 md1>
            <div class="caption grey--text">文件大小</div>
            <div>{{ (project.file_size / 1024 / 1024).toFixed(2) }}M</div>
          </v-flex>
          <v-flex class="pa-2" xs12 sm12 md2>
            <v-row justify="end">
              <v-btn-toggle shaped>
                <v-btn :href="project.apk_download_url">
                  <v-icon color="success">mdi-cloud-download</v-icon>
                </v-btn>
                <qrcode-dialog
                  :description="project.apk_description"
                  :qrcode="project.apk_download_url"
                ></qrcode-dialog>
                <confirm-dialog
                  @onDeleteFile="deleteFile({ id: project.id })"
                ></confirm-dialog>
              </v-btn-toggle>
            </v-row>
          </v-flex>
        </v-layout>
        <v-divider></v-divider>
      </v-card>
    </div>
    <div v-else>
      <v-card-text class="text-center"
        ><h3>列表空空如也~~赶紧上传文件看看吧</h3></v-card-text
      >
    </div>
  </v-col>
</template>

<script>
import QrcodeDialog from "@/components/QrcodeDialog";
import ConfirmDialog from "@/components/ConfirmDialog";
import { mapGetters, mapActions } from "vuex";

export default {
  components: {
    QrcodeDialog,
    ConfirmDialog,
  },
  computed: {
    ...mapGetters(["sortFiles"]),
  },
  methods: {
    ...mapActions(["fetchFileList", "deleteFile"]),
  },
  mounted() {
    this.fetchFileList();
  },
};
</script>

<style>
#download {
  width: 95%;
  padding: 15px;
  margin: 40px auto 40px;
  font-size: 14px;
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.4);
}
.project.complete {
  border-left: 4px solid #3cd1c2;
}
</style>
