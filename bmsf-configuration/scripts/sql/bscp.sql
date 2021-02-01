CREATE DATABASE IF NOT EXISTS bscpdb;

USE bscpdb;

CREATE TABLE IF NOT EXISTS `t_sharding` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fkey` varchar(64) NOT NULL,
  `Fdb_id` varchar(64) NOT NULL,
  `Fdb_name` varchar(64) NOT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `idx_t_sharding_key` (`Fkey`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_sharding_db` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fdb_id` varchar(64) NOT NULL,
  `Fhost` varchar(64) NOT NULL,
  `Fport` int(11) DEFAULT NULL,
  `Fuser` varchar(32) NOT NULL,
  `Fpassword` varchar(32) NOT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `idx_t_sharding_db_db_id` (`Fdb_id`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_proc_attr` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fcloud_id` varchar(64) NOT NULL,
  `Fip` varchar(32) NOT NULL,
  `Fbiz_id` varchar(64) NOT NULL,
  `Fapp_id` varchar(64) NOT NULL,
  `Fpath` varchar(512) NOT NULL,
  `Flabels` longtext,
  `Fcreator` varchar(64) NOT NULL,
  `Flast_modify_by` varchar(64) DEFAULT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `uidx_attr` (`Fcloud_id`,`Fip`,`Fbiz_id`,`Fapp_id`,`Fpath`),
  KEY `idx_state` (`Fstate`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_system` (
    `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `Fcurrent_version` varchar(64) NOT NULL,
    `Fkind` varchar(64) NOT NULL,
    `Foperator` varchar(64) DEFAULT NULL,
    `Fcreate_time` datetime(3) DEFAULT NULL,
    `Fupdate_time` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`Fid`),
    UNIQUE KEY `idx_kind` (`Fkind`),
    KEY `idx_cversion` (`Fcurrent_version`),
    KEY `idx_operator` (`Foperator`),
    KEY `idx_ctime` (`Fcreate_time`),
    KEY `idx_utime` (`Fupdate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `t_system`(`Fcurrent_version`, `Fkind`, `Foperator`, `Fcreate_time`, `Fupdate_time`) SELECT 'v0.0.0-202011201517', '__PRODUCT_KIND__', 'system-init', NOW(), NOW() FROM DUAL WHERE NOT EXISTS (SELECT * FROM `t_system`);

CREATE TABLE IF NOT EXISTS `t_local_auth` (
    `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `Fp_type` varchar(64) NOT NULL,
    `Fv0` varchar(64) NOT NULL,
    `Fv1` varchar(64) NOT NULL,
    `Fv2` varchar(64) NOT NULL,
    `Fv3` varchar(64) NOT NULL,
    `Fv4` varchar(64) NOT NULL,
    `Fv5` varchar(64) NOT NULL,
    PRIMARY KEY (`Fid`),
    UNIQUE KEY `uidx_policy` (`Fp_type`,`Fv0`,`Fv1`,`Fv2`,`Fv3`,`Fv4`,`Fv5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE DATABASE IF NOT EXISTS bscp_default;

USE bscp_default;

CREATE TABLE IF NOT EXISTS `t_app_instance` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fbiz_id` varchar(64) NOT NULL,
  `Fapp_id` varchar(64) NOT NULL,
  `Fcloud_id` varchar(64) NOT NULL,
  `Fip` varchar(32) NOT NULL,
  `Fpath` varchar(512) NOT NULL,
  `Flabels` longtext,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `idx_unionids` (`Fbiz_id`,`Fapp_id`,`Fcloud_id`,`Fip`,`Fpath`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_app_instance_release` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Finstance_id` bigint(20) NOT NULL,
  `Fbiz_id` varchar(64) NOT NULL,
  `Fapp_id` varchar(64) NOT NULL,
  `Fcfg_id` varchar(64) NOT NULL,
  `Frelease_id` varchar(64) NOT NULL,
  `Feffect_time` datetime(3) DEFAULT NULL,
  `Feffect_code` int(11) DEFAULT NULL,
  `Feffect_msg` varchar(128) DEFAULT NULL,
  `Freload_time` datetime(3) DEFAULT NULL,
  `Freload_code` int(11) DEFAULT NULL,
  `Freload_msg` varchar(128) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `idx_unionids` (`Finstance_id`,`Fcfg_id`,`Frelease_id`),
  KEY `idx_effected` (`Fcfg_id`,`Frelease_id`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_application` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fapp_id` varchar(64) NOT NULL,
  `Fbiz_id` varchar(64) NOT NULL,
  `Fname` varchar(64) NOT NULL,
  `Fdeploy_type` int(11) DEFAULT NULL,
  `Fcreator` varchar(64) NOT NULL,
  `Flast_modify_by` varchar(64) DEFAULT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `idx_t_application_app_id` (`Fapp_id`),
  UNIQUE KEY `idx_bizidname` (`Fbiz_id`,`Fname`),
  KEY `idx_state` (`Fstate`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_audit` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fsource_type` int(11) DEFAULT NULL,
  `Fop_type` int(11) DEFAULT NULL,
  `Fbiz_id` varchar(64) NOT NULL,
  `Fsource_id` varchar(64) NOT NULL,
  `Foperator` varchar(64) NOT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  KEY `idx_operator` (`Foperator`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`),
  KEY `idx_sourcetype` (`Fsource_type`),
  KEY `idx_optype` (`Fop_type`),
  KEY `idx_bid` (`Fbiz_id`),
  KEY `idx_sourceid` (`Fsource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_commit` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fcommit_id` varchar(64) NOT NULL,
  `Fbiz_id` varchar(64) NOT NULL,
  `Fapp_id` varchar(64) NOT NULL,
  `Fcfg_id` varchar(64) NOT NULL,
  `Fcommit_mode` int(11) DEFAULT NULL,
  `Foperator` varchar(64) NOT NULL,
  `Frelease_id` varchar(64) NOT NULL,
  `Fmulti_commit_id` varchar(64) NOT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `uidx_multi` (`Fcfg_id`,`Fmulti_commit_id`),
  UNIQUE KEY `idx_t_commit_commit_id` (`Fcommit_id`),
  KEY `idx_operator` (`Foperator`),
  KEY `idx_releaseid` (`Frelease_id`),
  KEY `idx_state` (`Fstate`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_bizid` (`Fbiz_id`),
  KEY `idx_appid` (`Fapp_id`),
  KEY `idx_mcommitid` (`Fmulti_commit_id`),
  KEY `idx_utime` (`Fupdate_time`),
  KEY `idx_cfgid` (`Fcfg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_content` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fbiz_id` varchar(64) NOT NULL,
  `Fapp_id` varchar(64) NOT NULL,
  `Fcfg_id` varchar(64) NOT NULL,
  `Fcommit_id` varchar(64) NOT NULL,
  `Findex` longtext,
  `Fcontent_id` varchar(64) NOT NULL,
  `Fcontent_size` bigint(20) unsigned DEFAULT NULL,
  `Fcreator` varchar(64) NOT NULL,
  `Flast_modify_by` varchar(64) DEFAULT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  KEY `idx_unionids` (`Fbiz_id`,`Fapp_id`,`Fcfg_id`,`Fcommit_id`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_config` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fcfg_id` varchar(64) NOT NULL,
  `Fbiz_id` varchar(64) NOT NULL,
  `Fapp_id` varchar(64) NOT NULL,
  `Fname` varchar(64) NOT NULL,
  `Ffpath` varchar(128) NOT NULL,
  `Fuser` varchar(64) NOT NULL,
  `Fuser_group` varchar(64) NOT NULL,
  `Ffile_privilege` varchar(64) NOT NULL,
  `Ffile_format` varchar(64) NOT NULL,
  `Ffile_mode` int(11) DEFAULT NULL,
  `Fcreator` varchar(64) NOT NULL,
  `Flast_modify_by` varchar(64) DEFAULT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `idx_t_config_cfg_id` (`Fcfg_id`),
  UNIQUE KEY `idx_appidnamepath` (`Fapp_id`,`Fname`,`Ffpath`),
  KEY `idx_bizid` (`Fbiz_id`),
  KEY `idx_state` (`Fstate`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_multi_commit` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fmulti_commit_id` varchar(64) NOT NULL,
  `Fbiz_id` varchar(64) NOT NULL,
  `Fapp_id` varchar(64) NOT NULL,
  `Foperator` varchar(64) NOT NULL,
  `Fmulti_release_id` varchar(64) NOT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `idx_t_multi_commit_multi_commit_id` (`Fmulti_commit_id`),
  KEY `idx_releaseid` (`Fmulti_release_id`),
  KEY `idx_state` (`Fstate`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`),
  KEY `idx_bid` (`Fbiz_id`),
  KEY `idx_appid` (`Fapp_id`),
  KEY `idx_operator` (`Foperator`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_multi_release` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fmulti_release_id` varchar(64) NOT NULL,
  `Fname` varchar(64) NOT NULL,
  `Fbiz_id` varchar(64) NOT NULL,
  `Fapp_id` varchar(64) NOT NULL,
  `Fstrategy_id` varchar(64) NOT NULL,
  `Fstrategies` longtext,
  `Fcreator` varchar(64) NOT NULL,
  `Fmulti_commit_id` varchar(64) NOT NULL,
  `Flast_modify_by` varchar(64) DEFAULT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `idx_t_multi_release_multi_release_id` (`Fmulti_release_id`),
  KEY `idx_name` (`Fname`),
  KEY `idx_bid` (`Fbiz_id`),
  KEY `idx_appid` (`Fapp_id`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_release` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Frelease_id` varchar(64) NOT NULL,
  `Fname` varchar(64) NOT NULL,
  `Fbiz_id` varchar(64) NOT NULL,
  `Fapp_id` varchar(64) NOT NULL,
  `Fcfg_id` varchar(64) NOT NULL,
  `Fcfg_name` varchar(64) NOT NULL,
  `Fcfg_fpath` varchar(128) NOT NULL,
  `Fuser` varchar(64) NOT NULL,
  `Fuser_group` varchar(64) NOT NULL,
  `Ffile_privilege` varchar(64) NOT NULL,
  `Ffile_format` varchar(64) NOT NULL,
  `Ffile_mode` int(11) DEFAULT NULL,
  `Fstrategy_id` varchar(64) NOT NULL,
  `Fstrategies` longtext,
  `Fcreator` varchar(64) NOT NULL,
  `Fcommit_id` varchar(64) NOT NULL,
  `Fmulti_release_id` varchar(64) DEFAULT NULL,
  `Flast_modify_by` varchar(64) DEFAULT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `idx_t_release_release_id` (`Frelease_id`),
  KEY `idx_mreleaseid` (`Fmulti_release_id`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`),
  KEY `idx_name` (`Fname`),
  KEY `idx_bid` (`Fbiz_id`),
  KEY `idx_appid` (`Fapp_id`),
  KEY `idx_cfgid` (`Fcfg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_strategy` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fstrategy_id` varchar(64) NOT NULL,
  `Fbiz_id` varchar(64) NOT NULL,
  `Fapp_id` varchar(64) NOT NULL,
  `Fname` varchar(64) NOT NULL,
  `Fcontent` longtext NOT NULL,
  `Fcreator` varchar(64) NOT NULL,
  `Flast_modify_by` varchar(64) DEFAULT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `idx_t_strategy_strategy_id` (`Fstrategy_id`),
  UNIQUE KEY `uidx_strategy` (`Fapp_id`,`Fname`),
  KEY `idx_bid` (`Fbiz_id`),
  KEY `idx_utime` (`Fupdate_time`),
  KEY `idx_state` (`Fstate`),
  KEY `idx_ctime` (`Fcreate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_template` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Ftemplate_id` varchar(64) NOT NULL,
  `Fbiz_id` varchar(64) NOT NULL,
  `Fname` varchar(64) NOT NULL,
  `Ffile_name` varchar(64) NOT NULL,
  `Ffile_path` varchar(128) NOT NULL,
  `Fuser` varchar(64) NOT NULL,
  `Fuser_group` varchar(64) NOT NULL,
  `Ffile_privilege` varchar(64) NOT NULL,
  `Ffile_format` varchar(64) NOT NULL,
  `Ffile_mode` int(11) DEFAULT NULL,
  `Fengine_type` int(11) DEFAULT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fcreator` varchar(64) NOT NULL,
  `Flast_modify_by` varchar(64) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `idx_t_template_template_id` (`Ftemplate_id`),
  UNIQUE KEY `idx_bizidname` (`Fbiz_id`,`Fname`),
  KEY `idx_utime` (`Fupdate_time`),
  KEY `idx_bizid` (`Fbiz_id`),
  KEY `idx_ctime` (`Fcreate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_template_bind` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fbiz_id` varchar(64) NOT NULL,
  `Ftemplate_id` varchar(64) NOT NULL,
  `Fapp_id` varchar(64) NOT NULL,
  `Fcfg_id` varchar(64) NOT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreator` varchar(64) NOT NULL,
  `Flast_modify_by` varchar(64) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `uidx_bind` (`Ftemplate_id`,`Fapp_id`),
  UNIQUE KEY `idx_t_template_bind_cfg_id` (`Fcfg_id`),
  KEY `idx_appid` (`Fapp_id`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`),
  KEY `idx_bizid` (`Fbiz_id`),
  KEY `idx_tplid` (`Ftemplate_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_template_version` (
  `Fid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Fversion_id` varchar(64) NOT NULL,
  `Fbiz_id` varchar(64) NOT NULL,
  `Ftemplate_id` varchar(64) NOT NULL,
  `Fversion_tag` varchar(64) NOT NULL,
  `Fcontent_id` varchar(64) NOT NULL,
  `Fcontent_size` bigint(20) unsigned DEFAULT NULL,
  `Fmemo` varchar(128) DEFAULT NULL,
  `Fcreator` varchar(64) NOT NULL,
  `Flast_modify_by` varchar(64) DEFAULT NULL,
  `Fstate` int(11) DEFAULT NULL,
  `Fcreate_time` datetime(3) DEFAULT NULL,
  `Fupdate_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`Fid`),
  UNIQUE KEY `idx_t_template_version_version_id` (`Fversion_id`),
  UNIQUE KEY `idx_version` (`Ftemplate_id`, `Fversion_tag`),
  KEY `idx_bizid` (`Fbiz_id`),
  KEY `idx_ctime` (`Fcreate_time`),
  KEY `idx_utime` (`Fupdate_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
