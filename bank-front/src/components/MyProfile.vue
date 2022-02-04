<template>
  <v-container>
    <v-row>
      <v-col>
        <v-data-table
          :headers="this.headers"
          :items="this.account.creditCards"
          class="elevation-1"
        >
          <template v-slot:[`item.accountNumber`]="">
            {{ showSecrets ? account.accountNumber : "****************" }}
          </template>
          <template v-slot:[`item.amount`]="">
            {{ account.amount }}
          </template>
          <template v-slot:[`item.secret`]="">
            {{ showSecrets ? account.secret : "**********" }}
          </template>
          <template v-slot:[`item.pan`]="{ item }">
            {{ showSecrets ? item.pan : "****************" }}
          </template>
          <template v-slot:[`item.cvc`]="{ item }">
            {{ showSecrets ? item.cvc : "***" }}
          </template>
          <template v-slot:top>
            <v-toolbar flat>
              <v-dialog v-model="dialogConfirmPassword" max-width="500px">
                <v-card>
                  <v-card-title>
                    <span class="text-h5">Confirm password</span>
                  </v-card-title>
                  <v-card-text>
                    <v-container>
                      <v-row>
                        <v-col cols="12" sm="8" md="12">
                          <v-text-field
                            type="password"
                            v-model="password"
                            label="Password"
                          ></v-text-field>
                        </v-col>
                      </v-row>
                    </v-container>
                  </v-card-text>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn
                      color="blue darken-1"
                      text
                      @click="closeConfirmPassword()"
                    >
                      Cancel
                    </v-btn>
                    <v-btn
                      color="blue darken-1"
                      text
                      @click="confirmPassword()"
                    >
                      Confirm
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </v-toolbar>
          </template>
          <template v-slot:[`item.actions`]="">
            <v-icon small @click="enableSecrets()" v-if="!showSecrets">
              mdi-eye
            </v-icon>
            <v-icon small @click="disableSecrets()" v-if="showSecrets">
              mdi-eye-off
            </v-icon>
          </template>
        </v-data-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";
import * as comm from "../configuration/communication.js";
export default {
  data() {
    return {
      account: {
        accountNumber: "",
        amount: 0,
        secret: "",
        creditCards: [
          {
            pan: "",
            cvc: "",
            holderName: "",
            validUntil: "",
          },
        ],
      },
      headers: [
        { text: "Account number", value: "accountNumber" },
        { text: "Amount", value: "amount" },
        { text: "Secret", value: "secret" },
        { text: "Pan", value: "pan" },
        { text: "Cvc", value: "cvc" },
        { text: "Holder name", value: "holderName" },
        { text: "Valid until", value: "validUntil" },
        { text: "Actions", value: "actions" },
      ],
      showSecrets: false,
      dialogConfirmPassword: false,
      password: "",
    };
  },
  mounted() {
    axios({
      method: "get",
      url: comm.BankProtocol + "://" + comm.BankServer + "/api/account",
      headers: comm.getHeader(),
    }).then((response) => {
      if (response.status == 200) {
        this.account = response.data;
      }
    });
  },
  methods: {
    enableSecrets() {
      this.dialogConfirmPassword = true;
    },
    disableSecrets() {
      this.showSecrets = false;
    },
    closeConfirmPassword() {
      this.dialogConfirmPassword = false;
    },
    confirmPassword() {
      axios({
        method: "post",
        url:
          comm.BankProtocol + "://" + comm.BankServer + "/api/confirm-password",
        headers: comm.getHeader(),
        data: JSON.stringify(this.password),
      })
        .then((response) => {
          console.log(response.data);
          this.dialogConfirmPassword = false;
          this.showSecrets = true;
          this.password = "";
        })
        .catch(() => {
          alert("Wrong password!");
          this.password = "";
        });
    },
  },
};
</script>
