<template>
  <v-form ref="form" v-model="valid" lazy-validation>
    <v-container>
      <v-row align="center" justify="center">
        <v-col cols="12" sm="4">
          <v-text-field
            v-model="username"
            :rules="[rules.required]"
            label="Username:"
            required
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row align="center" justify="center">
        <v-col cols="12" sm="4">
          <v-text-field
            v-model="password"
            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            :rules="[rules.required]"
            :type="showPassword ? 'text' : 'password'"
            label="Password"
            @click:append="showPassword = !showPassword"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row align="center" justify="center">
        <v-col cols="12" sm="4">
          <v-btn :disabled="!valid" color="success" class="mr-4" @click="login">
            Log in
          </v-btn>
        </v-col>
      </v-row>
      <v-row align="center" justify="center">
        <v-col cols="12" sm="4">
          <v-alert outlined dense type="error" v-model="alert">
            Username or password is incorrect
          </v-alert>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<script>
import axios from "axios";
import * as comm from "../configuration/communication.js";
import * as validator from "../plugins/validator.js";
import eventBus from "../plugins/eventBus.js"
export default {
  data() {
    return {
      showPassword: false,
      showPassCode: false,
      valid: true,
      username: "",
      password: "",
      rules: validator.rules,
      alert: false,
    };
  },
  mounted() {},
  methods: {
    login() {
      this.alert = false;
      if (this.$refs.form.validate()) {
        let credentials = {
          username: this.username,
          password: this.password,
        };
        axios({
          method: "post",
          url: comm.BankProtocol + "://" + comm.BankServer + "/api/login",
          data: JSON.stringify(credentials),
        })
          .then((response) => {
            if (response.status == 200) {
              comm.setJWTToken(response.data);
              eventBus.$emit('login');
              this.$router.push({ name: "MyProfile" });
            }
          })
          .catch(() => {
            this.alert = true;
          });
      }
    },
  },
};
</script>
