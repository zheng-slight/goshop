/*
 Navicat Premium Data Transfer

 Source Server         : 本机数据库
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : ginshoprbac

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 10/10/2025 20:45:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for access
-- ----------------------------
DROP TABLE IF EXISTS `access`;
CREATE TABLE `access`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `module_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0',
  `type` tinyint(1) NULL DEFAULT NULL COMMENT '类型: 1 表示模块, 2, 表示菜单: 显示在模块下面的栏位, 3 表示操作',
  `action_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单, 操作url',
  `module_id` int(0) NULL DEFAULT NULL,
  `sort` int(0) NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  `status` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 114 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of access
-- ----------------------------
INSERT INTO `access` VALUES (52, '管理员管理', 1, '', '', 0, 1001, '管理员管理', 0, 1);
INSERT INTO `access` VALUES (53, '角色管理', 1, '', '', 0, 100, '角色管理', 0, 1);
INSERT INTO `access` VALUES (54, '管理员管理', 2, '管理员列表', 'manager', 52, 100, '管理员列表', 0, 1);
INSERT INTO `access` VALUES (55, '管理员管理', 2, '增加管理员', 'manager/add', 52, 100, '管理员列表', 0, 1);
INSERT INTO `access` VALUES (56, '管理员管理', 3, '编辑管理员', 'manager/edit', 52, 100, '编辑管理员', 0, 1);
INSERT INTO `access` VALUES (57, '管理员管理', 3, '删除管理员', 'manager/delete', 52, 12, '删除管理员', 0, 1);
INSERT INTO `access` VALUES (59, '角色管理', 2, '角色列表', 'role', 53, 100, '角色列表', 0, 1);
INSERT INTO `access` VALUES (60, '角色管理', 2, '增加角色', 'role/add', 53, 1001, '增加角色', 0, 1);
INSERT INTO `access` VALUES (61, '角色管理', 3, '编辑角色', 'role/edit', 53, 100, '编辑角色', 0, 1);
INSERT INTO `access` VALUES (62, '角色管理', 3, '删除角色', 'role/delete', 53, 100, '删除角色', 0, 1);
INSERT INTO `access` VALUES (63, '权限管理', 1, '', '', 0, 100, '权限管理', 0, 1);
INSERT INTO `access` VALUES (64, '权限管理', 2, '权限列表', 'access', 63, 100, '', 0, 1);
INSERT INTO `access` VALUES (67, '权限管理', 2, '增加权限', 'access/add', 63, 100, '', 0, 1);
INSERT INTO `access` VALUES (68, '轮播图管理', 1, '', '', 0, 100, '', 0, 1);
INSERT INTO `access` VALUES (69, '轮播图管理', 2, '轮播图列表', 'focus', 68, 101, '1111', 0, 1);
INSERT INTO `access` VALUES (70, '轮播图管理', 2, '增加轮播图', 'focus/add', 68, 100, '增加轮播图', 0, 1);
INSERT INTO `access` VALUES (71, '轮播图管理', 3, '编辑轮播图', 'focus/edit', 68, 100, '', 0, 1);
INSERT INTO `access` VALUES (76, '管理员管理', 3, '执行增加管理员', 'manager/doAdd', 52, 100, '执行增加', 0, 1);
INSERT INTO `access` VALUES (77, '管理员管理', 3, '执行修改管理员', 'manager/doEdit', 52, 100, '执行修改', 0, 1);
INSERT INTO `access` VALUES (78, '角色管理', 3, '执行增加角色', 'role/doAdd', 53, 100, '执行增加', 0, 1);
INSERT INTO `access` VALUES (79, '角色管理', 3, '执行修改角色', 'role/doEdit', 53, 100, '执行修改', 0, 1);
INSERT INTO `access` VALUES (80, '角色管理', 3, '角色授权', 'role/auth', 53, 100, '', 0, 1);
INSERT INTO `access` VALUES (81, '角色管理', 3, '执行角色授权', 'role/doAuth', 53, 100, '执行授权', 0, 1);
INSERT INTO `access` VALUES (82, '权限管理', 3, '修改权限', 'access/edit', 63, 100, '执行修改', 0, 1);
INSERT INTO `access` VALUES (83, '权限管理', 3, '删除权限', 'access/delete', 63, 100, '', 0, 1);
INSERT INTO `access` VALUES (84, '权限管理', 3, '执行增加权限', 'access/doAdd', 63, 100, '', 0, 1);
INSERT INTO `access` VALUES (85, '权限管理', 3, '执行修改权限', 'access/doEdit', 63, 100, '执行修改\r\n', 0, 1);
INSERT INTO `access` VALUES (88, '轮播图管理', 3, '执行增加轮播图', 'focus/doAdd', 68, 100, '', 0, 1);
INSERT INTO `access` VALUES (89, '轮播图管理', 3, '执行修改轮播图', 'focus/doEdit', 68, 100, '', 0, 1);
INSERT INTO `access` VALUES (90, '轮播图管理', 3, '删除轮播图', 'focus/delete', 68, 100, '', 0, 1);
INSERT INTO `access` VALUES (91, '商品管理', 1, '', '', 0, 100, '', 0, 1);
INSERT INTO `access` VALUES (92, '商品管理', 2, '商品分类', 'goodsCate', 91, 100, '', 0, 1);
INSERT INTO `access` VALUES (93, '商品管理', 3, '增加商品分类', 'goodsCate/add', 91, 100, '', 0, 1);
INSERT INTO `access` VALUES (94, '商品管理', 3, '执行增加商品分类', 'goodsCate/doAdd', 91, 100, '', 0, 1);
INSERT INTO `access` VALUES (95, '商品管理', 3, '修改商品分类', 'goodsCate/edit', 91, 100, '', 0, 1);
INSERT INTO `access` VALUES (96, '商品管理', 3, '执行修改商品分类', 'goodsCate/doEdit', 91, 100, '', 0, 1);
INSERT INTO `access` VALUES (97, '商品管理', 3, '删除商品分类', 'goodsCate/delete', 91, 100, '', 0, 1);
INSERT INTO `access` VALUES (98, '商品管理', 2, '商品类型', 'goodsType', 91, 100, '', 0, 1);
INSERT INTO `access` VALUES (99, '商品管理', 3, '增加商品类型', 'goodsType/add', 91, 100, '', 0, 1);
INSERT INTO `access` VALUES (100, '商品管理', 3, '执行增加商品类型', 'goodsType/doAdd', 91, 100, '', 0, 1);
INSERT INTO `access` VALUES (101, '商品管理', 3, '删除商品类型', 'goodsType/delete', 91, 100, '', 0, 1);
INSERT INTO `access` VALUES (102, '商品管理', 3, '修改商品类型', 'goodsType/edit', 91, 100, '', 0, 1);
INSERT INTO `access` VALUES (103, '商品管理', 3, '执行修改商品类型', 'goodsType/doEdit', 91, 100, '', 0, 1);
INSERT INTO `access` VALUES (104, '商品管理', 2, '商品列表', 'goods', 91, 100, '', 0, 1);
INSERT INTO `access` VALUES (105, '导航管理', 1, '', '', 0, 100, '', 0, 1);
INSERT INTO `access` VALUES (106, '导航管理', 2, '导航列表', 'nav/index', 106, 100, '', 0, 1);
INSERT INTO `access` VALUES (107, '导航管理', 2, '导航列表', 'nav', 105, 100, '', 0, 1);
INSERT INTO `access` VALUES (108, '导航管理', 3, '增加导航', 'nav/add', 105, 100, '', 0, 1);
INSERT INTO `access` VALUES (109, '导航管理', 3, '执行增加导航', 'nav/doAdd', 105, 100, '', 0, 1);
INSERT INTO `access` VALUES (110, '导航管理', 3, '删除导航', 'nav/delete', 105, 100, '', 0, 1);
INSERT INTO `access` VALUES (111, '商品管理', 3, '执行修改导航', 'nav/doEdit', 105, 100, '', 0, 1);
INSERT INTO `access` VALUES (112, '导航管理', 3, '执修导航', 'nav/edit', 105, 100, '', 0, 1);
INSERT INTO `access` VALUES (113, '系统设置', 1, '', '', 0, 100, '', 0, 1);
INSERT INTO `access` VALUES (114, '系统设置', 2, '设置列表', 'setting', 113, 100, '', 0, 1);

-- ----------------------------
-- Table structure for manager
-- ----------------------------
DROP TABLE IF EXISTS `manager`;
CREATE TABLE `manager`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `mobile` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  `role_id` int(0) NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  `is_super` tinyint(1) NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `role_id`(`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of manager: admin 默认密码: 123456
-- ----------------------------
INSERT INTO `manager` VALUES (1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '152016111', '5188611114@qq.com', 1, 9, 0, 1);
INSERT INTO `manager` VALUES (2, 'zhangsan', 'e10adc3949ba59abbe56e057f20f883e', '1520111122', '342338691122@qq.com', 1, 14, 1581661532, 0);
INSERT INTO `manager` VALUES (6, 'lisi', 'e10adc3949ba59abbe56e057f20f883e', '1520171111', '11114292@qq.com', 1, 16, 1631156378, 0);
INSERT INTO `manager` VALUES (9, '测试', 'e10adc3949ba59abbe56e057f20f883e', '', '', 1, 4, 1677398426, 0);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (4, '编辑部门', '这是一个编辑部门，主要负责文章编辑', 1, 1591062092);
INSERT INTO `role` VALUES (14, '软件部门', '软件部门', 0, 1631075350);
INSERT INTO `role` VALUES (17, '运营部门', '', 1, 1690706958);

-- ----------------------------
-- Table structure for role_access
-- ----------------------------
DROP TABLE IF EXISTS `role_access`;
CREATE TABLE `role_access`  (
  `role_id` int(0) NOT NULL,
  `access_id` int(0) NOT NULL,
  INDEX `role_id`(`role_id`) USING BTREE,
  INDEX `access_id`(`access_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_access
-- ----------------------------
INSERT INTO `role_access` VALUES (4, 52);
INSERT INTO `role_access` VALUES (4, 54);
INSERT INTO `role_access` VALUES (4, 53);
INSERT INTO `role_access` VALUES (4, 59);
INSERT INTO `role_access` VALUES (4, 63);
INSERT INTO `role_access` VALUES (4, 64);
INSERT INTO `role_access` VALUES (4, 67);
INSERT INTO `role_access` VALUES (4, 83);
INSERT INTO `role_access` VALUES (4, 105);
INSERT INTO `role_access` VALUES (4, 107);
INSERT INTO `role_access` VALUES (4, 108);
INSERT INTO `role_access` VALUES (4, 113);
INSERT INTO `role_access` VALUES (4, 114);
INSERT INTO `role_access` VALUES (14, 63);
INSERT INTO `role_access` VALUES (14, 64);
INSERT INTO `role_access` VALUES (14, 68);
INSERT INTO `role_access` VALUES (14, 69);
INSERT INTO `role_access` VALUES (14, 113);
INSERT INTO `role_access` VALUES (14, 114);

SET FOREIGN_KEY_CHECKS = 1;
