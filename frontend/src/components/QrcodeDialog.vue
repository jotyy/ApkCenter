<template>
  <v-dialog width="300" height="400">
    <template v-slot:activator="{ on }">
      <v-btn v-on="on">
        <v-icon color="info">mdi-qrcode</v-icon>
      </v-btn>
    </template>

    <v-card>
      <v-container>
        <v-row class="mt-1 justify-center align-center"><v-card-subtitle>{{description}}</v-card-subtitle></v-row>
        <v-row class="justify-center align-center">
          <qrcode
            :value="qrcode"
            :options="{ width: 250, height: 250 }"
          ></qrcode>
        </v-row>
      </v-container>
     
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  props: ["description", "qrcode"],
  data() {
    return {
      dialog: false,
    };
  },
};
</script>
