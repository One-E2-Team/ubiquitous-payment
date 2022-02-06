<template>
  <v-container>
    <v-alert
      v-if="alert"
      :value="alertText"
      color="red"
      type="error"
      dismissible
      text
      v-model="alertText"
      >{{ alertText }}</v-alert
    >
    <v-row align="center" justify="center">
      <v-col cols="12" sm="9">
        <v-form ref="form1" lazy-validation>
          <v-row align="center" justify="center">
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="credentials.username"
                label="Username:"
                required
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row align="center" justify="center">
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="credentials.name"
                label="Name:"
                required
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row align="center" justify="center">
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="credentials.email"
                :rules="[rules.email, rules.required]"
                label="Mail:"
                required
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row align="center" justify="center">
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="credentials.password"
                :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
                :rules="[rules.password]"
                :type="show ? 'text' : 'password'"
                label="Password"
                hint="At least 8 characters, 1 lower, 1 capital letter, 1 number and 1 special character"
                counter
                @click:append="show = !show"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row align="center" justify="center">
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="password2"
                :rules="[passwordMatch, rules.required]"
                :type="'password'"
                label="Confirm password"
                hint="Password must match"
                counter
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row align="center" justify="center">
            <v-col cols="12" sm="6">
              <v-radio-group v-model="credentials.role" mandatory row>
                <v-radio label="Customer" value="CUSTOMER"></v-radio>
                <v-radio label="Seller" value="SELLER"></v-radio>
              </v-radio-group>
            </v-col>
          </v-row>
        </v-form>
        <v-row align="center" justify="center">
          <v-col cols="12" sm="6" class="d-flex justify-space-around mb-6">
            <v-btn color="primary" @click="register"> Register </v-btn>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";
import * as comm from "../configuration/communication.js";
import * as validator from "../plugins/validator.js";
export default {
  data() {
    return {
      alert: false,
      alertText: "",
      show: false,
      credentials: {
        username: "",
        email: "",
        password: "",
        name: "",
        role: "",
      },
      password2: "",
      rules: validator.rules,
      passwordMatch: () =>
        this.credentials.password === this.password2 || "Password must match",
    };
  },
  mounted() {},
  methods: {
    register() {
      let data = {
        username: this.credentials.username,
        email: this.credentials.email,
        password: this.credentials.password,
        name: this.credentials.name,
        role: this.credentials.role,
      };
      axios({
        method: "post",
        url: comm.WSprotocol + "://" + comm.WSserver + "/api/users",
        data: JSON.stringify(data),
      }).then((response) => {
        if (response.status == 200) {
          alert("Check your email!");
          this.alert = false;
          if (response.data.message == "Invalid data.") {
            this.alert = true;
            if (response.data.errors.includes("Password")) {
              this.alertText =
                "Password is too weak. Please choose another password.";
            }
          }
          if (response.data.message == "Server error while registering.") {
            this.alert = true;
            this.alertText =
              "Chosen e-mail already exists.Please choose another mail.";
          }
        } else if (response.status == 201) {
          this.$router.push({ name: "Login" });
        }
      });
    },
  },
};
</script>
