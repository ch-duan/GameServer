CREATE TABLE `game`.`user`  (
                                    `uuid` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                                    `phone` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                                    `openid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                    `nike_name` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                                    `status` int(10) NOT NULL,
                                    `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                    `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    `modify_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
                                    `deleted` tinyint(1) NOT NULL,
                                    PRIMARY KEY (`uuid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;