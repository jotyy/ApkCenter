<template>
  <v-dialog v-model="loading" width="200" height="150" persistent>
    <v-card width="200" height="150">
      <v-container>
        <v-row class="fill-height" align-content="center" justify="center">
          <v-col class="subtitle-1 text-center" cols="12">
            正在上传，请稍候..
          </v-col>
          <v-col cols="6">
            <v-progress-linear
              color="secondary"
              indeterminate
              rounded
              height="6"
            ></v-progress-linear>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </v-dialog>
</template>

<script>
import { mapState } from "vuex";

export default {
  computed: {
    ...mapState(["loading"]),
  },
};
</script>
