<template>
  <div>psadjasdjdsaasd</div>
</template>

<script>
import axios from "axios";
import * as comm from "../configuration/communication.js";
import * as validator from "../plugins/validator.js";
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
  mounted() {
    axios({
      method: "get",
      url: comm.BankProtocol + "://" + comm.BankServer + "/api/account",
      headers: comm.getHeader(),
    }).then((response) => {
      if (response.status == 200) {
        console.log(response);
      }
    });
  },
  methods: {},
};
</script>
