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
  >{{alertText}}</v-alert>
  <v-row align="center" justify="center">
      <v-col cols="auto">
          <h2>Register your Web-shop:</h2>
      </v-col>
  </v-row>
  <v-row align="center" justify="center">
    <v-col cols="12" sm="9" >
        <v-form ref="form1" lazy-validation>
             <v-row align="center" justify="center">
              <v-col cols="12" sm="6" >
                <v-text-field
                    v-model="credentials.username"
                    label="Username:"
                    :rules="[ rules.required ]"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row align="center" justify="center">
              <v-col cols="12" sm="6" >
                <v-text-field
                    v-model="credentials.webShopName"
                    label="Webshop name:"
                    :rules="[ rules.required ]"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row align="center" justify="center">
              <v-col cols="12" sm="6">
                <v-text-field
                    v-model="credentials.password"
                    :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
                    :rules="[ rules.password ]"
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
        </v-form>
          <v-row align="center" justify="center">
              <v-col cols="12" sm="6" class="d-flex justify-space-around mb-6">
                <v-btn
                color="primary"
                @click="register">
                Register webshop
                </v-btn>
              </v-col>
          </v-row>
      </v-col>
    </v-row>
  </v-container>

</template>

<script>
import axios from 'axios'
import * as comm from '../configuration/communication.js'
import * as validator from '../plugins/validator.js'
  export default {
    data() {return {
      alert: false,
      alertText : '',
      show: false,
      credentials: {
        username : '',
        password: '',
        webShopName: '',
        paymentTypes: []
      },
      password2: '',
      rules: validator.registrationRules,
      passwordMatch: () => (this.credentials.password === this.password2) || 'Password must match',
    }},
    mounted(){
    },
    methods: {
      register() {
          let data = {
            username : this.credentials.username,  
            password: this.credentials.password,
            webShopName : this.credentials.webShopName,
            paymentTypes : this.paymentTypes
           }
          axios({
            method: "post",
            url: comm.Protocol + "://" + comm.PSPserver +"/api/psp/register-web-shop",
            data: JSON.stringify(data),
          }).then((response) => {
            if (response.status == 200) {
                this.credentials = {}
                alert("Your web-shop is successfully registered.");
                this.$router.push({name: "Home"});
            if(response.data.message=="Invalid data."){
                this.alert = true;
                  if(response.data.errors.includes("Password")){
                    this.alertText = "Password is too weak. Please choose another password."
                  }
            }
            if(response.data.message == "Server error while registering."){
                    this.alert = true;
                    this.alertText = "Error while registering."
                }
            }
          })
        }
      }
  }
</script>