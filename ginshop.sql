/*
 Navicat Premium Data Transfer

 Source Server         : 本机数据库
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : ginshop

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 10/10/2025 20:44:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `uid` int(0) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `default_address` tinyint(1) NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 47 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of address
-- ----------------------------
INSERT INTO `address` VALUES (42, 12, '李四', '15201686411', '北京市 海淀区 西二旗 xxx好11', 0, 0);
INSERT INTO `address` VALUES (43, 12, '张三', '15201686412', '深圳市   宝安区  xxx 2222222', 0, 0);
INSERT INTO `address` VALUES (44, 12, '王五', '15201686412', '上海市 xxx11 222 111', 0, 0);
INSERT INTO `address` VALUES (45, 12, '赵六', '15201686412', '发顺丰', 0, 0);
INSERT INTO `address` VALUES (46, 12, '毛七', '15555555555', '成都市武侯区', 1, 1683726543);

-- ----------------------------
-- Table structure for focus
-- ----------------------------
DROP TABLE IF EXISTS `focus`;
CREATE TABLE `focus`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `focus_type` tinyint(1) NULL DEFAULT NULL,
  `focus_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `link` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` int(0) NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of focus
-- ----------------------------
INSERT INTO `focus` VALUES (14, '小E手机', 1, 'static/upload/20230411/1681220459500617100.png', 'http://www.xxx.com', 1001, 0, 1631677671);
INSERT INTO `focus` VALUES (15, '2022北京车展招展函下载', 1, '', 'http://www.xxx.com', 100, 1, 1631677692);
INSERT INTO `focus` VALUES (16, '小E电视1', 2, 'static/upload/20230301/1677676829.jpg', 'http://www.xxx.com', 100, 0, 1631679244);
INSERT INTO `focus` VALUES (17, '测', 1, 'static/upload/20230228/1677591510.png', 'www.baidu.com', 1003, 0, 1677591510);
INSERT INTO `focus` VALUES (18, 'df', 1, 'static/upload/20230411/1681220429445020000.png', 'www.baidu.com', 100, 1, 1677591574);
INSERT INTO `focus` VALUES (21, '鼠标', 0, '', '', 10, 1, 1677765620);
INSERT INTO `focus` VALUES (22, '鼠标', 0, '', 'www.baidu.com', 10, 1, 1677765740);

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `sub_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `goods_sn` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `cate_id` int(0) NULL DEFAULT NULL,
  `click_count` int(0) NULL DEFAULT NULL,
  `goods_number` int(0) NULL DEFAULT NULL,
  `price` decimal(10, 2) NULL DEFAULT NULL,
  `market_price` decimal(10, 2) NULL DEFAULT NULL,
  `relation_goods` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `goods_attr` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `goods_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `goods_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `goods_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `goods_gift` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `goods_fitting` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `goods_keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `goods_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `goods_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `is_delete` tinyint(0) NULL DEFAULT NULL,
  `is_hot` tinyint(0) NULL DEFAULT NULL,
  `is_best` tinyint(0) NULL DEFAULT NULL,
  `is_new` tinyint(0) NULL DEFAULT NULL,
  `goods_type_id` int(0) NULL DEFAULT NULL,
  `sort` int(0) NULL DEFAULT NULL,
  `status` tinyint(0) NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES (35, '华为Plus', '华为Plus大法师的大法师的大法师的大法师的', '', 23, 100, 11, 163.21, 111.33, '11', '颜色:红色,白色,黄色 | 尺寸:41,42,43', '1,2', '1', 'static/upload/20230423/1682255187294832300.png', '11', '111', '第三大神大法师的大法师的大法师的大法师的', '大法师的大法师的大法师的大法师的大法师的大法师的', '<p><img src=\"http://xxx-xxx-1312546108.cos.ap-chengdu.myqcloud.com/static/upload/20230408/1680958106694877300.png\" style=\"width: 300px;\" class=\"fr-fic fr-dib\">发顺丰乳沟如果大幅改善的风格</p><p><br></p><p><br></p>', 0, 1, 1, 1, 10, 1, 1, 1678370701);
INSERT INTO `goods` VALUES (36, '华为Plus1', '华为Plus', '', 23, 100, 11, 163.21, 111.33, '11', '是多少', '1,2', '1', '', '11', '111', '第三大神', '大法师的', '发顺丰乳沟如果大幅改善的风格', 0, 1, 0, 1, 10, 1, 1, 1678546468);
INSERT INTO `goods` VALUES (37, '华为Plus111', '华为Plus', '', 23, 100, 11, 163.21, 111.33, '11', '是多少', '1,2', '1', '', '11', '111', '第三大神', '大法师的', '<p>发顺丰乳沟如果大幅改善的风格</p>', 0, 1, 0, 1, 6, 1, 1, 1678546477);
INSERT INTO `goods` VALUES (38, '海尔冰箱1', '海尔冰箱', '', 30, 100, 0, 111.00, 111.00, '', '颜色:红色,白色,黄色 | 尺寸:41,42,43', '1,5', '1', 'static/upload/20230313/1678712778377732600.png', '', '', '', '', '<p>的撒发刚阿萨德噶的搜嘎是搭嘎是的</p>', 0, 1, 1, 0, 11, 0, 1, 1678712778);
INSERT INTO `goods` VALUES (39, '手机1', '手机1', '', 27, 100, 0, 111.00, 111.00, '35', '', '1,2', '1', 'static/upload/20230419/1681911034696613800.png', '36', '', '', '手机1', '<p>手机1</p>', 0, 0, 1, 0, 10, 0, 1, 1681911034);

-- ----------------------------
-- Table structure for goods_attr
-- ----------------------------
DROP TABLE IF EXISTS `goods_attr`;
CREATE TABLE `goods_attr`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `goods_id` int(0) NULL DEFAULT NULL,
  `attribute_cate_id` int(0) NULL DEFAULT NULL,
  `attribute_id` int(0) NULL DEFAULT NULL,
  `attribute_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `attribute_type` tinyint(1) NULL DEFAULT NULL,
  `attribute_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `sort` int(0) NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 107 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods_attr
-- ----------------------------
INSERT INTO `goods_attr` VALUES (4, 36, 10, 8, '尺寸', 1, '11x343', 10, 1678546468, 1);
INSERT INTO `goods_attr` VALUES (5, 36, 10, 11, '是否支持蓝牙', 3, '是\r\n', 10, 1678546468, 1);
INSERT INTO `goods_attr` VALUES (6, 36, 10, 12, '颜色', 3, '红\r\n', 10, 1678546468, 1);
INSERT INTO `goods_attr` VALUES (29, 38, 11, 18, '门类', 3, '单开门\r\n', 10, 1678799660, 1);
INSERT INTO `goods_attr` VALUES (30, 38, 11, 19, '节能等级', 3, '一级节能\r\n', 10, 1678799660, 1);
INSERT INTO `goods_attr` VALUES (31, 37, 6, 13, '内存', 1, '8G', 10, 1678799670, 1);
INSERT INTO `goods_attr` VALUES (32, 37, 6, 14, '硬盘', 1, '8G', 10, 1678799670, 1);
INSERT INTO `goods_attr` VALUES (33, 37, 6, 15, '显示器', 1, '11x343', 10, 1678799671, 1);
INSERT INTO `goods_attr` VALUES (34, 37, 6, 16, '支持蓝牙', 3, '是\r\n', 10, 1678799671, 1);
INSERT INTO `goods_attr` VALUES (35, 37, 6, 17, '性能', 2, '不交货', 10, 1678799671, 1);
INSERT INTO `goods_attr` VALUES (81, 39, 10, 8, '尺寸', 1, '11x343', 10, 1681911066, 1);
INSERT INTO `goods_attr` VALUES (82, 39, 10, 11, '是否支持蓝牙', 3, '是\r\n', 10, 1681911066, 1);
INSERT INTO `goods_attr` VALUES (83, 39, 10, 12, '颜色', 3, '红\r\n', 10, 1681911066, 1);
INSERT INTO `goods_attr` VALUES (99, 35, 10, 8, '尺寸', 1, '11x343', 10, 1682258609, 1);
INSERT INTO `goods_attr` VALUES (100, 35, 10, 11, '是否支持蓝牙', 3, '是\r\n', 10, 1682258609, 1);
INSERT INTO `goods_attr` VALUES (101, 35, 10, 12, '颜色', 3, '红\r\n', 10, 1682258609, 1);
INSERT INTO `goods_attr` VALUES (102, 35, 10, 20, '性能', 2, '### 第一代骁龙®8+移动平台\r\nSoC 工艺：台积电4nm工艺制程 \\n\r\nCPU 主频：八核处理器，最高主频可达：3.2GHz \\n\r\nGPU ：Adreno™ GPU 图形处理器 \\n\r\nAI：第七代 AI 引擎 \\n', 10, 1682258609, 1);
INSERT INTO `goods_attr` VALUES (103, 35, 10, 22, '内存与容量', 2, '## 12GB+512GB 最高可选\r\n运行内存：8GB / 12GB LPDDR5 高速内存（6400Mbps）\\n\r\n机身存储：256GB / 512GB UFS 3.1 高速存储 \\n', 10, 1682258609, 1);
INSERT INTO `goods_attr` VALUES (104, 35, 10, 23, '外观容量', 1, '长度：163.17mm 宽度：74.97mm 厚度：9.06mm 重量：225g', 10, 1682258609, 1);
INSERT INTO `goods_attr` VALUES (105, 35, 10, 24, '充电与电池', 2, '## 4860mAh(typ) / 4760mAh(min)\r\n内置单电芯高能量密度电池，不可拆卸 \\n\r\nUSB Type-C 双面充电接口 \\n\r\n手机支持 QC4 / QC3+ / QC3.0 / QC2.0 / PD3.0 / PD2.0 快充协议+MI FC 2.0 快充 \\n\r\n67W 小E澎湃秒充 / 50W 无线快充 / 10W 无线反充 \\n', 10, 1682258610, 1);
INSERT INTO `goods_attr` VALUES (106, 35, 10, 25, '影像系统', 1, 'MIUI 13', 10, 1682258610, 1);
INSERT INTO `goods_attr` VALUES (107, 35, 10, 26, '传感器', 1, '超声波距离传感器丨环境光传感器丨加速度传感器丨陀螺仪丨电子罗盘｜X 轴线性马达丨 红外线遥控器丨气压计丨后置光线（色温）传感器丨Flicker 传感器丨激光对焦传感器', 10, 1682258610, 1);

-- ----------------------------
-- Table structure for goods_cate
-- ----------------------------
DROP TABLE IF EXISTS `goods_cate`;
CREATE TABLE `goods_cate`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `cate_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `link` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `template` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pid` int(0) NULL DEFAULT NULL,
  `sub_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `keywords` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `description` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  `sort` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 37 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods_cate
-- ----------------------------
INSERT INTO `goods_cate` VALUES (23, '手机', '', '', '', 0, '', '', '手机', 1, '13', 1631938178);
INSERT INTO `goods_cate` VALUES (24, '电视', '', '', '', 0, '', '', '手机', 1, '10', 1631938196);
INSERT INTO `goods_cate` VALUES (25, '笔记本 平板', '', '', '', 0, '', '', '手机', 1, '11', 1631938209);
INSERT INTO `goods_cate` VALUES (26, '家电', '', '', '', 0, '', '', '手机', 1, '9', 1631938214);
INSERT INTO `goods_cate` VALUES (27, '小E11', 'static/upload/20230411/1681221296101196600.png', '', '', 23, '', '小E11', '小E手机官网正品小E11推荐，小E手机小E11最新价格，有多种颜色可选，另有小E11详细介绍及图片，还有', 1, '9', 1631938291);
INSERT INTO `goods_cate` VALUES (28, 'Redmi 11A', 'static/upload/20230411/1681221312291691900.png', 'http://www.xxx.com', 'bbbb.html', 23, '', '防疫', '游戏必备', 1, '8', 1631938339);
INSERT INTO `goods_cate` VALUES (29, '小E电视55寸', 'static/upload/20210918/1631938567.jpg', '', '', 24, '', '', '', 1, '10', 1631938567);
INSERT INTO `goods_cate` VALUES (30, '冰箱', 'static/upload/20210918/1631940993.jpg', 'http://www.xxx.com', '', 26, '', '', '', 1, '10', 1631938591);
INSERT INTO `goods_cate` VALUES (35, '鼠标', 'static/upload/20230302/1677765818.jpg', 'www.baidu.com', '鼠标', 0, '', '鼠标', '鼠标', 1, '111', 1677765818);
INSERT INTO `goods_cate` VALUES (36, '雷神鼠标', 'static/upload/20230302/1677766229.png', 'www.baidu.com', '', 35, '', '', '', 1, '102', 1677766229);

-- ----------------------------
-- Table structure for goods_color
-- ----------------------------
DROP TABLE IF EXISTS `goods_color`;
CREATE TABLE `goods_color`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `color_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `color_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `status` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods_color
-- ----------------------------
INSERT INTO `goods_color` VALUES (1, '红色', 'red', 1);
INSERT INTO `goods_color` VALUES (2, '黑色', '#000', 1);
INSERT INTO `goods_color` VALUES (3, '黄色', 'yellow', 1);
INSERT INTO `goods_color` VALUES (4, '金色', '#ebf10f', 1);
INSERT INTO `goods_color` VALUES (5, '灰色', '#eee', 1);

-- ----------------------------
-- Table structure for goods_image
-- ----------------------------
DROP TABLE IF EXISTS `goods_image`;
CREATE TABLE `goods_image`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `goods_id` int(0) NULL DEFAULT NULL,
  `img_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `color_id` int(0) NULL DEFAULT NULL,
  `sort` int(0) NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods_image
-- ----------------------------
INSERT INTO `goods_image` VALUES (3, 37, 'static/upload/20230311/1678546857222693200.png', 0, 10, 1678546858, 1);
INSERT INTO `goods_image` VALUES (4, 37, 'static/upload/20230311/1678546857218464700.jpg', 0, 10, 1678546858, 1);
INSERT INTO `goods_image` VALUES (10, 38, 'static/upload/20230313/1678712777170831700.png', 0, 10, 1678712778, 1);
INSERT INTO `goods_image` VALUES (11, 38, 'static/upload/20230313/1678712777166998300.jpg', 0, 10, 1678712778, 1);
INSERT INTO `goods_image` VALUES (16, 39, 'static/upload/20230419/1681911064616894700.jpg', 0, 10, 1681911066, 1);
INSERT INTO `goods_image` VALUES (17, 39, 'static/upload/20230419/1681911064620084400.png', 0, 10, 1681911066, 1);
INSERT INTO `goods_image` VALUES (18, 39, 'static/upload/20230419/1681911064663898300.jpg', 0, 10, 1681911066, 1);
INSERT INTO `goods_image` VALUES (19, 35, 'static/upload/20230423/1682255185173729700.png', 0, 10, 1682255187, 1);
INSERT INTO `goods_image` VALUES (20, 35, 'static/upload/20230423/1682255185156367200.png', 2, 10, 1682255187, 1);
INSERT INTO `goods_image` VALUES (21, 35, 'static/upload/20230423/1682255185163325500.png', 2, 10, 1682255187, 1);
INSERT INTO `goods_image` VALUES (22, 35, 'static/upload/20230423/1682255185513326600.png', 2, 10, 1682255187, 1);
INSERT INTO `goods_image` VALUES (23, 35, 'static/upload/20230423/1682255185549041900.png', 0, 10, 1682255188, 1);

-- ----------------------------
-- Table structure for goods_type
-- ----------------------------
DROP TABLE IF EXISTS `goods_type`;
CREATE TABLE `goods_type`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` int(0) NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods_type
-- ----------------------------
INSERT INTO `goods_type` VALUES (5, '电视', '电视', 0, 1632299512);
INSERT INTO `goods_type` VALUES (6, '电脑', '电脑', 1, 1632299526);
INSERT INTO `goods_type` VALUES (7, '路由器', '路由器', 1, 1632299535);
INSERT INTO `goods_type` VALUES (9, '衣服', '衣服', 0, 1632361292);
INSERT INTO `goods_type` VALUES (10, '手机', '手机\r\n', 1, 1677767983);
INSERT INTO `goods_type` VALUES (11, '冰箱', '冰箱一类家电', 1, 1678712824);

-- ----------------------------
-- Table structure for goods_type_attribute
-- ----------------------------
DROP TABLE IF EXISTS `goods_type_attribute`;
CREATE TABLE `goods_type_attribute`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `cate_id` int(0) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `attr_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `attr_value` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  `sort` int(0) NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `cate_id`(`cate_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods_type_attribute
-- ----------------------------
INSERT INTO `goods_type_attribute` VALUES (8, 10, '尺寸', '1', '', 1, 101, 1632370232);
INSERT INTO `goods_type_attribute` VALUES (9, 7, '颜色', '2', '', 1, 10, 1632370271);
INSERT INTO `goods_type_attribute` VALUES (11, 10, '是否支持蓝牙', '3', '是\r\n否', 1, 10, 1678003308);
INSERT INTO `goods_type_attribute` VALUES (12, 10, '颜色', '3', '红\r\n黄\r\n蓝\r\n黑\r\n白\r\n灰', 1, 10, 1678003400);
INSERT INTO `goods_type_attribute` VALUES (13, 6, '内存', '1', '', 0, 10, 1678197409);
INSERT INTO `goods_type_attribute` VALUES (14, 6, '硬盘', '1', '', 1, 10, 1678197419);
INSERT INTO `goods_type_attribute` VALUES (15, 6, '显示器', '1', '', 1, 10, 1678197428);
INSERT INTO `goods_type_attribute` VALUES (16, 6, '支持蓝牙', '3', '是\r\n否', 1, 10, 1678197511);
INSERT INTO `goods_type_attribute` VALUES (17, 6, '性能', '2', '', 1, 10, 1678197530);
INSERT INTO `goods_type_attribute` VALUES (18, 11, '门类', '3', '单开门\r\n双开门', 1, 10, 1678712874);
INSERT INTO `goods_type_attribute` VALUES (19, 11, '节能等级', '3', '一级节能\r\n二级节能\r\n三级节能', 1, 10, 1678712937);
INSERT INTO `goods_type_attribute` VALUES (20, 10, '性能', '2', '', 1, 10, 1682256049);
INSERT INTO `goods_type_attribute` VALUES (21, 10, '相机', '2', '', 1, 10, 1682256062);
INSERT INTO `goods_type_attribute` VALUES (22, 10, '内存与容量', '2', '', 1, 10, 1682256073);
INSERT INTO `goods_type_attribute` VALUES (23, 10, '外观容量', '1', '', 1, 10, 1682256090);
INSERT INTO `goods_type_attribute` VALUES (24, 10, '充电与电池', '2', '', 1, 10, 1682256126);
INSERT INTO `goods_type_attribute` VALUES (25, 10, '影像系统', '1', '', 1, 10, 1682256147);
INSERT INTO `goods_type_attribute` VALUES (26, 10, '传感器', '1', '', 1, 10, 1682256159);

-- ----------------------------
-- Table structure for nav
-- ----------------------------
DROP TABLE IF EXISTS `nav`;
CREATE TABLE `nav`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `position` tinyint(1) NULL DEFAULT NULL,
  `is_opennew` tinyint(1) NULL DEFAULT NULL,
  `relation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `sort` int(0) NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of nav
-- ----------------------------
INSERT INTO `nav` VALUES (1, '小E商城1', 'http://www.xxx.com', 2, 2, '36,35', 10, 1, 1592919226);
INSERT INTO `nav` VALUES (2, 'MIUI', 'http://www.xxx.com', 1, 1, '1', 10, 1, 1592921999);
INSERT INTO `nav` VALUES (3, '小E手机', 'https://shouji.mi.com/', 2, 2, '19,20', 10, 1, 1592922081);
INSERT INTO `nav` VALUES (4, '小E电视', 'https://ds.mi.com/', 2, 2, '23,24', 10, 1, 1592922273);
INSERT INTO `nav` VALUES (5, '路由器', 'http://bbs.xxx.com', 2, 1, '25', 10, 1, 1592922331);
INSERT INTO `nav` VALUES (8, '云服务', 'https://i.mi.com/', 1, 2, '2', 10, 1, 1593529309);
INSERT INTO `nav` VALUES (9, '金融', 'https://jr.mi.com/?from=micom', 1, 1, '1', 10, 1, 1593529329);
INSERT INTO `nav` VALUES (10, '有品', 'https://youpin.mi.com/', 1, 1, '1', 10, 1, 1593529346);
INSERT INTO `nav` VALUES (11, '家电', '', 2, 1, '1', 10, 1, 1593529451);
INSERT INTO `nav` VALUES (12, '智能电视', '', 2, 1, '1', 10, 1, 1593529470);
INSERT INTO `nav` VALUES (14, '小E帮助中心', 'http://www.xxx.com', 3, 2, '12,13,14', 101, 1, 1634788777);

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `uid` int(0) NULL DEFAULT NULL,
  `order_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `all_price` decimal(10, 2) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `zipcode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `pay_status` tinyint(1) NULL DEFAULT NULL,
  `pay_type` tinyint(1) NULL DEFAULT NULL,
  `order_status` tinyint(1) NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  `pay_time` int(0) NULL DEFAULT NULL,
  `distribution_time` int(0) NULL DEFAULT NULL,
  `exwarehouse_time` int(0) NULL DEFAULT NULL,
  `successful_time` int(0) NULL DEFAULT NULL,
  `cancel_time` int(0) NULL DEFAULT NULL,
  `return_time` int(0) NULL DEFAULT NULL,
  `logistics_company` int(0) NULL DEFAULT NULL,
  `waybill_no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 41 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES (39, 12, '202112161333074546', 4698.00, '李四', '15201686411', '北京市 海淀区 西二旗 xxx好11', NULL, 0, 0, 0, 1639632787, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `order` VALUES (40, 12, '202112161339577439', 4698.00, '王五', '15201686412', '上海市 xxx11 222 111', NULL, 0, 0, 0, 1639633197, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `order` VALUES (41, 12, '202305131649468546', 326.42, '毛七', '15555555555', '成都市武侯区', NULL, 0, 0, 0, 1683967786, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for order_item
-- ----------------------------
DROP TABLE IF EXISTS `order_item`;
CREATE TABLE `order_item`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `order_id` int(0) NULL DEFAULT NULL,
  `uid` int(0) NULL DEFAULT NULL,
  `product_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `product_id` int(0) NULL DEFAULT NULL,
  `product_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `product_price` decimal(10, 2) NULL DEFAULT NULL,
  `product_num` int(0) NULL DEFAULT NULL,
  `goods_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `goods_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 54 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_item
-- ----------------------------
INSERT INTO `order_item` VALUES (54, 41, 12, '华为Plus', 35, 'static/upload/20230423/1682255187294832300.png', 163.21, 2, '1', '', 0);

-- ----------------------------
-- Table structure for setting
-- ----------------------------
DROP TABLE IF EXISTS `setting`;
CREATE TABLE `setting`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `site_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `site_logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `site_keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `site_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `no_picture` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `site_icp` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `site_tel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `search_keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `tongji_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `appid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `app_secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `end_point` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `bucket_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `oss_status` tinyint(1) NULL DEFAULT NULL,
  `oss_domain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `thumbnail_size` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

 
-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `last_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  `status` tinyint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (12, 'e10adc3949ba59abbe56e057f20f883e', '19950326585', '127.0.0.1', '', 1683589268, 1);

-- ----------------------------
-- Table structure for user_temp
-- ----------------------------
DROP TABLE IF EXISTS `user_temp`;
CREATE TABLE `user_temp`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `send_count` int(0) NULL DEFAULT NULL,
  `add_day` int(0) NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  `sign` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 41 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_temp
-- ----------------------------
INSERT INTO `user_temp` VALUES (40, '127.0.0.1', '19950326585', 2, 20230508, 1683557661, 'fe7d2c80d53a4b89b78a3bf07e271688');
INSERT INTO `user_temp` VALUES (41, '127.0.0.1', '19950326585', 1, 20230509, 1683589161, 'c006085c940046308b932a7e4a908c0f');

SET FOREIGN_KEY_CHECKS = 1;
