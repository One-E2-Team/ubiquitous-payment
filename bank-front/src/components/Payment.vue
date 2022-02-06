<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="auto">
        <h2>
          In the following form, please enter your credit card information in
          order to complate transaction. Amount: {{ paymentDetails.amount }}
          {{ paymentDetails.currency }}
          {{
            paymentDetails.amount != paymentDetails.amountRsd
              ? "(" + paymentDetails.amountRsd + " RSD)"
              : ""
          }}
        </h2>
      </v-col>
    </v-row>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="9">
        <v-form ref="form" v-model="valid" lazy-validation>
          <v-row align="center" justify="center">
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="creditCard.pan"
                label="PAN:"
                placeholder="0000000000000000"
                :rules="[rules.required, rules.pan]"
                required
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row align="center" justify="center">
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="creditCard.validUntil"
                label="Valid until:"
                placeholder="01/22"
                :rules="[rules.required, rules.cardValid]"
                required
              ></v-text-field>
            </v-col>
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="creditCard.cvc"
                label="CVC:"
                placeholder="999"
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
                placeholder="John Doe"
                :rules="[rules.required]"
                required
              ></v-text-field>
            </v-col>
          </v-row>
        </v-form>
        <v-row align="center" justify="center">
          <v-col cols="12" sm="6" class="d-flex justify-space-around mb-6">
            <v-btn color="primary" @click="pay()"> Pay </v-btn>
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
  mounted() {
    this.loadPaymentDetails();
  },
  data() {
    return {
      valid: true,
      rules: validator.rules,
      paymentDetails: {
        amount: "",
        amountRsd: "",
        currency: "",
      },
      creditCard: {
        pan: "",
        cvc: "",
        holderName: "",
        validUntil: "",
      },
    };
  },
  methods: {
    loadPaymentDetails() {
      const urlId = comm.getUrlVars()["id"];
      axios({
        method: "get",
        url:
          comm.BankProtocol +
          "://" +
          comm.BankServer +
          "/api/payment-details/" +
          urlId,
      }).then((response) => {
        if (response.status == 200) {
          this.paymentDetails = response.data;
        }
      });
    },
    pay() {
      if (!this.$refs.form.validate()) {
        return;
      }
      const urlId = comm.getUrlVars()["id"];
      if (!urlId) {
        alert("Url id missing");
        return;
      }

      axios({
        method: "post",
        url: comm.BankProtocol + "://" + comm.BankServer + "/api/pay/" + urlId,
        data: JSON.stringify(this.creditCard),
      }).then((response) => {
        if (response.status == 200) {
          window.location.href = response.data;
        }
      });
    },
  },
};
</script>

<style>
</style>
