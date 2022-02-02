<template>
  <v-container>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="9">
        <v-form ref="form1" lazy-validation>
          <v-row align="center" justify="center">
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="creditCard.pan"
                label="PAN:"
                :rules="[rules.required]"
                required
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row align="center" justify="center">
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="creditCard.validUntil"
                label="Valid until:"
                :rules="[rules.required]"
                required
              ></v-text-field>
            </v-col>
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="creditCard.cvc"
                label="CVC:"
                :rules="[rules.required]"
                required
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row align="center" justify="center">
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="creditCard.holderName"
                label="Holder name:"
                :rules="[rules.required]"
                required
              ></v-text-field>
            </v-col>
          </v-row>
        </v-form>
        <v-row align="center" justify="center">
          <v-col cols="12" sm="6" class="d-flex justify-space-around mb-6">
            <v-btn color="primary" @click="pay()" :disabled="!isValid()">
              Pay
            </v-btn>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import * as validator from "../plugins/validator.js";
import axios from "axios";
import * as comm from "../configuration/communication.js";

export default {
  name: "Payment",
  mounted() {},
  data() {
    return {
      rules: validator.rules,
      paymentUrlId: "",
      creditCard: {
        pan: "",
        cvc: "",
        holderName: "",
        validUntil: "",
      },
    };
  },
  methods: {
    pay() {
      axios({
        method: "post",
        url:
          comm.BankProtocol +
          "://" +
          comm.BankServer +
          "/api/pay/" +
          comm.getUrlVars()["id"],
        data: JSON.stringify(this.creditCard),
      }).then((response) => {
        if (response.status == 200) {
          console.log(response);
          window.location.href = response.data;
        }
      });
    },
    isValid() {
      return (
        this.creditCard &&
        this.creditCard.pan &&
        this.creditCard.cvc &&
        this.creditCard.holderName &&
        this.creditCard.validUntil
      );
    },
  },
};
</script>

<style>
</style>
