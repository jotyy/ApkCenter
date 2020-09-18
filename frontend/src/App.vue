<template>
  <v-app>
    <v-app-bar app color="primary" light>
      <div class="d-flex align-center">
        <v-img
          alt="Apk Center Logo"
          class="shrink mr-2"
          contain
          src="./assets/android.png"
          transition="scale-transition"
          width="40"
        />

        <v-toolbar-title class="white--text weight-600"
          >ANDROID APK CENTER</v-toolbar-title
        >
      </div>

      <v-spacer></v-spacer>

      <v-btn href="https://jotyy.top/" target="_blank" text color="white">
        <span class="mr-2 white--text d-none d-md-flex">WRITTEN BY JOTYY</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
    </v-app-bar>

    <upload-dialog />
    <progress-dialog />

    <v-content>
      <Download />
      <!-- <v-snackbar v-model="snackbar" timeout="2">
        {{errMessage}}
      </v-snackbar> -->
    </v-content>
  </v-app>
</template>

<script>
import Download from "@/components/Download";
import UploadDialog from "@/components/UploadDialog";
import ProgressDialog from "@/components/ProgressDialog";

export default {
  name: "App",

  components: {
    Download,
    UploadDialog,
    ProgressDialog,
  },
};
</script>
