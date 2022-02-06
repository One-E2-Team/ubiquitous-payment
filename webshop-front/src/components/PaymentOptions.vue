<template>
  <v-container>
    <v-row>
      <v-col>
        <v-select
          @input="getAccounts"
          :items="this.paymentNames"
          label="Select payment type"
          persistent-hint
          return-object
          single-line
        ></v-select>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-data-table
          :headers="this.headers"
          :items="this.accounts"
          class="elevation-1"
        >
          <template v-slot:[`item.accountId`]="">
            {{ "**************************" }}
          </template>
          <template v-slot:[`item.secret`]="">
            {{ "**************************" }}
          </template>
          <template v-slot:top>
            <v-toolbar flat>
              <v-toolbar-title>{{ selectedPaymentOption }}</v-toolbar-title>
              <v-divider class="mx-4" inset vertical></v-divider>
              <v-spacer></v-spacer>
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
              <v-dialog v-model="dialog" max-width="500px">
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                    color="primary"
                    dark
                    class="mb-2"
                    v-bind="attrs"
                    v-on="on"
                    >New Account</v-btn
                  >
                </template>
                <v-card>
                  <v-card-title>
                    <span class="text-h5">{{ formTitle }}</span>
                  </v-card-title>

                  <v-card-text>
                    <v-container>
                      <v-row>
                        <v-col cols="12" sm="8" md="12">
                          <v-text-field
                            v-model="editedItem.accountId"
                            label="Account ID"
                          ></v-text-field>
                        </v-col>
                      </v-row>
                      <v-row>
                        <v-col cols="12" sm="6" md="12">
                          <v-text-field
                            v-model="editedItem.secret"
                            label="Secret"
                          ></v-text-field>
                        </v-col>
                      </v-row>
                    </v-container>
                  </v-card-text>

                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="close">
                      Cancel
                    </v-btn>
                    <v-btn color="blue darken-1" text @click="createAccount()">
                      Save
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
              <v-dialog v-model="dialogDelete" max-width="500px">
                <v-card>
                  <v-card-title class="text-h5"
                    >Are you sure you want to delete this item?</v-card-title
                  >
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="closeDelete"
                      >Cancel</v-btn
                    >
                    <v-btn color="blue darken-1" text @click="deleteItemConfirm"
                      >OK</v-btn
                    >
                    <v-spacer></v-spacer>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </v-toolbar>
          </template>
          <template v-slot:[`item.actions`]="{ item }">
            <v-icon small class="mr-2" @click="editItem(item)">
              mdi-pencil
            </v-icon>
            <v-icon small @click="deleteItem(item)"> mdi-delete </v-icon>
          </template>
          <template v-slot:no-data>
            <v-btn color="primary" @click="initialize">Reset</v-btn>
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
      paymentOptions: [],
      selectedPaymentOption: "",
      selectedPaymentOptionId: 0,
      paymentNames: [],
      dialogConfirmPassword: false,
      dialog: false,
      dialogDelete: false,
      editMode: false,
      password: "",
      headers: [
        {
          text: "Accounts",
          align: "start",
          sortable: false,
          value: "name",
        },
        { text: "ID", value: "ID" },
        { text: "Account ID", value: "accountId" },
        { text: "Secret", value: "secret" },
        { text: "Actions", value: "actions" },
      ],
      accounts: [],
      editedIndex: -1,
      editedItem: {
        ID: "",
        accountId: "",
        secret: "",
      },
      defaultItem: {
        ID: "",
        accountId: "",
        secret: "",
      },
    };
  },
  created() {
    axios({
      method: "get",
      url: comm.WSprotocol + "://" + comm.WSserver + "/api/payment-types",
      headers: comm.getHeader(),
    })
      .then((response) => {
        if (response.status == 200) {
          this.paymentOptions = response.data;
          for (let el of this.paymentOptions) {
            this.paymentNames.push(el.name);
          }
        }
      })
      .catch(() => {
        console.log("error");
      });
  },
  computed: {
    formTitle() {
      return !this.editMode ? "Add a new account" : "Edit your account";
    },
  },
  watch: {
    dialog(val) {
      val || this.close();
    },
    dialogDelete(val) {
      val || this.closeDelete();
    },
  },
  methods: {
    getAccounts(paymentType) {
      this.selectedPaymentOption = paymentType;
      for (let i = 0; i < this.paymentOptions.length; i++) {
        if (this.paymentOptions[i].name == paymentType) {
          this.selectedPaymentOptionId = this.paymentOptions[i].ID;
          break;
        }
      }
      axios({
        method: "get",
        url:
          comm.WSprotocol +
          "://" +
          comm.WSserver +
          "/api/accounts/" +
          paymentType,
        headers: comm.getHeader(),
      })
        .then((response) => {
          if (response.status == 200) {
            this.accounts = response.data;
            for (let account of this.accounts) {
              account.accountId = account.accountId.Data;
              account.secret = account.secret.Data;
            }
            this.selectedPaymentOptionId = response.data[0].paymentTypeId;
          }
        })
        .catch(() => {
          console.log("error");
        });
    },
    createAccount() {
      let data = {
        paymentTypeId: this.selectedPaymentOptionId,
        accountId: this.editedItem.accountId,
        secret: this.editedItem.secret,
      };
      if (this.editMode) {
        console.log("Editing..");
        axios({
          method: "put",
          url:
            comm.WSprotocol +
            "://" +
            comm.WSserver +
            "/api/accounts/" +
            this.editedItem.ID,
          data: JSON.stringify(data),
          headers: comm.getHeader(),
        })
          .then((response) => {
            console.log(response.status);
            this.getAccounts(this.selectedPaymentOption);
          })
          .catch(() => {
            console.log("error");
          });
        this.close();
      } else {
        if (this.accounts.length > 0) {
          alert(
            "You cannot have more than one account for the same payment type!"
          );
        } else {
          axios({
            method: "post",
            url: comm.WSprotocol + "://" + comm.WSserver + "/api/accounts",
            data: JSON.stringify(data),
            headers: comm.getHeader(),
          })
            .then((response) => {
              console.log(response.status);
              this.getAccounts(this.selectedPaymentOption);
            })
            .catch(() => {
              console.log("error");
            });
        }
        this.close();
      }
      this.editMode = false;
    },
    close() {
      this.dialog = false;
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem);
        this.editMode = false;
      });
    },
    deleteItem(item) {
      this.editedItem = Object.assign({}, item);
      this.dialogDelete = true;
    },

    deleteItemConfirm() {
      axios({
        method: "delete",
        url:
          comm.WSprotocol +
          "://" +
          comm.WSserver +
          "/api/accounts/" +
          this.editedItem.ID,
        headers: comm.getHeader(),
      })
        .then((response) => {
          console.log(response.status);
          this.getAccounts(this.selectedPaymentOption);
        })
        .catch(() => {
          console.log("error");
        });
      this.closeDelete();
      this.accounts = [];
    },

    editItem(item) {
      this.editedItem = Object.assign({}, item);
      this.dialogConfirmPassword = true;
      this.editMode = true;
    },
    closeConfirmPassword() {
      this.dialogConfirmPassword = false;
    },
    closeDelete() {
      this.dialogDelete = false;
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem);
        this.editedIndex = -1;
      });
    },
    confirmPassword() {
      axios({
        method: "put",
        url: comm.WSprotocol + "://" + comm.WSserver + "/api/confirm-password",
        headers: comm.getHeader(),
        data: JSON.stringify(this.password),
      })
        .then((response) => {
          console.log(response.data);
          this.dialogConfirmPassword = false;
          this.dialog = true;
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
