<template>
  <v-container>
    <v-row>
      <v-col>
        <v-data-table
          :headers="this.transactionHeaders"
          :items="this.transactions"
          class="elevation-1"
        >
          <template v-slot:top>
            <v-toolbar flat>
              <v-toolbar-title>All transactions</v-toolbar-title>
              <v-divider class="mx-4" inset vertical></v-divider>
              <v-spacer></v-spacer>
            </v-toolbar>
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
      transactions: [
        {
          amount: "",
          currency: "",
          acquirerAccountNumber: "",
          issuerPan: "",
          timestamp: "",
          transactionStatus: "",
        },
      ],
      transactionHeaders: [
        { text: "Amount", value: "amount" },
        { text: "Currency", value: "currency" },
        { text: "Acquirer account number", value: "acquirerAccountNumber" },
        { text: "Issuer pan", value: "issuerPan" },
        { text: "Timestamp", value: "timestamp" },
        { text: "Transaction status", value: "transactionStatus" },
      ],
    };
  },
  mounted() {
    this.getAllTransactions();
  },
  methods: {
    getAllTransactions() {
      axios({
        method: "get",
        url:
          comm.BankProtocol + "://" + comm.BankServer + "/api/all-transactions",
        headers: comm.getHeader(),
      }).then((response) => {
        if (response.status == 200) {
          this.transactions = response.data;
        }
      });
    },
  },
};
</script>
