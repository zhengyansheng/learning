<template>
  <div class="container">
    <div class="left">

      <h2>1) v-html</h2>
      <span> {{ website }} </span>
      <br>
      <span v-html="website"> </span>
      <hr>

      <h2>2) JavaScript 表达式</h2>
      {{ number+1 }}
      <br>
      {{ ok }}
      <br>
      {{ ok? "YES": "NO" }}
      <br>
      {{ message.split("").reverse().join(", ") }}
      <br>
      <h4>{{ fruitData.message }}</h4>
      <p>水果的总价是: {{ fruitData.price * fruitData.total}}</p>
      <hr>

      <h2>3) v-on 内置指令</h2>
      <button v-on:click="number+=1">加1</button>
      <button v-on:click="number-=1">减1</button>
      <br>
      <span>{{ number}}</span>
      <br>
      <button v-on:click="say('hello')">古诗</button>
      <hr>

      <h2>4) v-text指令 用来更新元素的文本内容</h2>
      <p>古诗欣赏: {{ message1 }}</p>
      <p v-text="message1"></p>

      <hr>
      <h2>5) v-once指令 只渲染元素和组件一次 </h2>
      <p v-once>不可改变: {{message2}}</p>
      <p>可以改变: {{message2}}</p>
      <p>
        <input type="text" v-model="message2" name="input2">
      </p>
      <hr>

      <h2>6) v-pre指令 不渲染 显示原始标签</h2>
      <p v-pre> {{ message }}</p>

      <h2>7) v-if指令 为true显示，否则不显示 </h2>

      <p v-if="ok1">西红柿</p>
      <p v-if="no">蔬菜</p>
      <p v-if="num>1000">蔬菜的库存很充足</p>
      <hr>

      <h2>x) 练习</h2>
      姓: <input type="text" v-model="firstName"><br>
      名称: <input type="text" v-model="lastName">
      <p>{{firstName}} {{lastName}}</p>


    </div>

    <div class="divider"></div>

    <div class="right">



      <h2>8) v-if/v-else-if指令 </h2>
      <span v-if="score>=90">优秀</span>
      <span v-else-if="score>=80">合格</span>
      <span v-else-if="score>=60">及格</span>
      <span v-else>不及格</span>
      <br>
      <input type="text" v-model="score">

      <hr>

      <h2>9) v-show </h2>
      <p v-show="ok1">西红柿</p>
      <p v-show="no">蔬菜</p>
      <p v-show="num>1000">蔬菜的库存很充足</p>

      <button v-on:click="switchOK()">点击我</button>
      <hr>

      <h2>10) v-for (push | unshift | filter) </h2>

      名称: <input type="text" v-model="deviceName"  placeholder="请输入"><br>
      产地: <input type="text" v-model="deviceCity"  placeholder="请输入"><br>
      价格: <input type="text" v-model="devicePrice"  placeholder="请输入"><br>
      <button @click="addvFor">添加</button>

      <ul>
        <li v-for="(item,index) in vForData" :key="item.name">
          {{index}}-{{item.name}}--{{item.city}}--{{item.price}}元
        </li>
      </ul>

      <p>所有商品</p>
      <ul>
        <li v-for="item in nameList" :key="item.name">{{item.name}} | {{item.city}} | {{item.price}}</li>
      </ul>

      <p>产地为上海的产品</p>
      <ul>
        <li v-for="item in nameLists()" :key="item.name">{{item.name}} | {{item.city}} | {{item.price}}</li>
      </ul>

      <p>价格大于或等于5000的商品</p>
      <ul>
        <li v-for="item in prices()" :key="item.name">{{item.name}} | {{item.city}} | {{item.price}}</li>
      </ul>

      <hr>
      <h2>11) v-for 加 v-if </h2>

      <p>已经出库的产品</p>
      <ul>
        <template v-for="item in goodss" :key="item.name">
          <li v-if="item.isOut">{{item.name}}</li>
        </template>
      </ul>

      <p>没有出库的产品</p>
      <ul>
        <template v-for="item in goodss" :key="item.name">
          <li v-if="!item.isOut">{{item.name}}</li>
        </template>
      </ul>

    </div>

  </div>


</template>

<script setup>

import {ref, reactive} from "vue";

// v-html
const website = ref("")
website.value = "<a href=\"https://www.baidu.com\">百度</a>"

// JavaScript表达式

const number = ref(1)
const ok = ref(true)
const message = ref("12345")
const fruitData = reactive({
  message: "fruit",
  price: 5,
  total: 260,
})

// v-on
const say = (param) => {
  alert("古诗来了......"+param)
}

// v-text
const message1 = ref()
message1.value = "白蛇转....."

// v-once
const message2 = ref("苹果")

// v-if
const ok1 = ref(true)
const no = ref(false)
const num = ref(1001)

// v-if-else
const score = ref(0)

const switchOK = () => {
  ok1.value = !ok1.value
}

// v-for
const vForData = reactive([
  {name:"洗衣机", city: "上海", price: "8600"},
  {name:"冰箱", city: "北京", price: "6800"},
  {name:"空调", city: "广州", price: "5900"},
])

// push
const deviceName = ref("")
const deviceCity = ref("")
const devicePrice = ref("")
const addvFor = () => {
  // vForData.push({name:deviceName.value, city: deviceCity.value, price: devicePrice.value})
  vForData.unshift({name:deviceName.value, city: deviceCity.value, price: devicePrice.value})
}

const nameList = reactive([
  {name:"洗衣机", city: "上海", price: "8600"},
  {name:"冰箱", city: "北京", price: "6800"},
  {name:"空调", city: "广州", price: "5900"},
  {name:"电视机", city: "上海", price: "4900"},
])

// onMounted(() => {
//   console.log('组件已挂载')
//   // 在这里可以执行例如获取数据、设置事件监听等操作
//   nameLists()
// })

const nameLists = () => {
    return nameList.filter(function (nameList) {
      return nameList.city === "上海";
    })
}

const prices = () => {
  return nameList.filter(function (nameList) {
    return nameList.price >= 5000;
  })
}

// v-for + v-if
const goodss = reactive([
  {name: "洗衣机", isOut: false},
  {name: "冰箱", isOut: true},
  {name: "空调", isOut: false},
  {name: "电视机", isOut: true},
  {name: "电脑", isOut: false},
])

// v-model
const firstName = ref("")
const lastName = ref("")



</script>

<style scoped>
.container {
  display: flex;
  height: 100vh; /* 页面高度 */
}

.left {
  flex: 1; /* 左侧占据 50% */
  //background-color: lightblue;
}

.right {
  padding-left: 20px;
  flex: 1; /* 右侧占据 50% */
  //background-color: lightgreen;
}

/* 竖线 */
.divider {
  width: 2px;
  background-color: black; /* 竖线颜色 */
}
</style>