<template>
    <div>
        <template>
          <v-container class="fill-height">
            <v-row>
              <v-col v-for="p in products" :key="p.id" cols="12" sm="4">
                <v-card class="mx-auto my-12" width="330" > 
                  <template slot="progress">
                    <v-progress-linear color="deep-purple" height="10" indeterminate ></v-progress-linear>
                  </template>
                    <!-- <v-img contain width="330" height="440" :src="protocol + '://' + server + '/data/' + p.picturePath"></v-img>
                    -->
                    <v-img contain width="330" height="340" :src="nesto"></v-img>
                  <v-card-title>{{p.name}}</v-card-title>

                  <v-card-text>
                    <div class="my-4 text-subtitle-1">
                      Price: {{p.price}} {{p.currency}}
                    </div>

                    <div>Description: {{p.description}}</div>
                  </v-card-text>

                  <v-divider class="mx-4"></v-divider>

                  <v-card-text>
                     <v-btn color="deep-purple lighten-2" text @click="makeOrder(p)" >
                            Order
                    </v-btn>
                  </v-card-text>
                </v-card>
              </v-col>
            </v-row>
          </v-container>
        </template>
    </div>
</template>

<script>
import axios from 'axios'
import * as comm from '../configuration/communication.js'
  export default {
    name: "HomePage",
    data() {return {
      products: [],
    }},
    mounted(){
       this.getProducts();
    },
    methods: {
     getProducts(){
       axios({
                method: "get",
                url: comm.WSprotocol +'://' + comm.WSserver + '/api/products',
            }).then(response => {
              if(response.status==200){
                this.products = response.data;
              }
            }).catch((response) => {
              console.log(response.data)
            });
     }, 
     makeOrder(product){
       axios({
                method: "post",
                url: comm.WSprotocol +'://' + comm.WSserver + '/api/orders/' + product.id,
                headers: comm.getHeader()
            }).then(response => {
              if(response.status==200){
                console.log("ok");
              }
            }).catch(() => {
              console.log("error")
            })
     }
    },
  }
</script>