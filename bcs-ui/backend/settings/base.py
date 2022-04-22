# -*- coding: utf-8 -*-
"""
Tencent is pleased to support the open source community by making 蓝鲸智云PaaS平台社区版 (BlueKing PaaS Community
Edition) available.
Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.

Django settings for backend project.

Generated by 'django-admin startproject' using Django 1.11.5.

For more information on this file, see
https://docs.djangoproject.com/en/1.11/topics/settings/

For the full list of settings and their values, see
https://docs.djangoproject.com/en/1.11/ref/settings/
"""

import os

# Apply pymysql patch
import pymysql
from django.conf.global_settings import gettext_noop as _

from backend.resources.utils import dynamic

pymysql.install_as_MySQLdb()

dynamic.patch_resource()

# Build paths inside the project like this: os.path.join(BASE_DIR, ...)
BASE_DIR = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

# Quick-start development settings - unsuitable for production
# See https://docs.djangoproject.com/en/1.11/howto/deployment/checklist/

# SECURITY WARNING: keep the secret key used in production secret!
SECRET_KEY = "jllc(^rzpe8_udv)oadny2j3ym#qd^x^3ns11_8kq(1rf8qpd2"

# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = True

ALLOWED_HOSTS = ["localhost", "127.0.0.1"]

INSTALLED_APPS = [
    "django.contrib.admin",
    "django.contrib.auth",
    "django.contrib.contenttypes",
    "django.contrib.sessions",
    "django.contrib.messages",
    "whitenoise.runserver_nostatic",
    "django.contrib.staticfiles",
    "rest_framework",
    "corsheaders",
    "django_extensions",
    "backend.accounts",
    "backend.bcs_web.audit_log.apps.AuditLogConfig",
    "backend.container_service.projects",
    "backend.container_service.misc.depot",
    "backend.container_service.clusters.apps.ClusterConfig",
    "backend.templatesets.legacy_apps.configuration.apps.TemplatesetsConfConfig",
    "backend.templatesets.apps.TemplateSetsConfig",
    "backend.templatesets.legacy_apps.instance.apps.TemplatesetsInstanceConfig",
    "backend.uniapps.resource",
    "backend.uniapps.network",
    "backend.helm.app",
    "backend.helm.helm",
    "backend.helm.authtoken.apps.HelmAuthtokenConfig",
    "backend.container_service.infras.hosts.terraform",
    # 模板集功能模块
    "backend.templatesets.var_mgmt.apps.VarMgmtConfig",
    "backend.change_log",
    "django_prometheus",
]

MIDDLEWARE = [
    # 该中间件必须在最前
    "django_prometheus.middleware.PrometheusBeforeMiddleware",
    "backend.accounts.middlewares.RequestProvider",
    "corsheaders.middleware.CorsMiddleware",
    "django.middleware.security.SecurityMiddleware",
    "backend.bcs_web.middleware.MultiDomainSessionMiddleware",
    "django.middleware.locale.LocaleMiddleware",
    "django.middleware.common.CommonMiddleware",
    "backend.bcs_web.middleware.MultiDomainCsrfViewMiddleware",
    "django.contrib.auth.middleware.AuthenticationMiddleware",
    "django.contrib.messages.middleware.MessageMiddleware",
    "django.middleware.clickjacking.XFrameOptionsMiddleware",
    # admin static file
    "whitenoise.middleware.WhiteNoiseMiddleware",
    # 该中间件必须在最后
    "django_prometheus.middleware.PrometheusAfterMiddleware",
]

# 是否启用 Django-Prometheus Migration，镜像构建时候禁用，原因是没有 DB 服务
PROMETHEUS_EXPORT_MIGRATIONS = os.environ.get('PROMETHEUS_EXPORT_MIGRATIONS', 'True') == 'True'

STATICFILES_STORAGE = "whitenoise.storage.CompressedManifestStaticFilesStorage"

ROOT_URLCONF = "backend.urls"

TEMPLATES = [
    {
        "BACKEND": "django.template.backends.django.DjangoTemplates",
        "DIRS": [
            os.path.join(BASE_DIR, "templates"),
            os.path.join(BASE_DIR, "backend/web_console/templates"),
            os.path.join(BASE_DIR, "backend/static"),
            os.path.join(BASE_DIR, "staticfiles"),
            os.path.join(BASE_DIR, "backend/templatesets/legacy_apps/configuration/yaml_mode/manifests"),
            os.path.join(BASE_DIR, "backend/helm/toolkit/kubehelm/templates"),
        ],
        "APP_DIRS": True,
        "OPTIONS": {
            "string_if_invalid": "{{%s}}",
            "context_processors": [
                "django.template.context_processors.debug",
                "django.template.context_processors.request",
                "django.contrib.auth.context_processors.auth",
                "django.contrib.messages.context_processors.messages",
                "backend.apps.context_processors.global_settings",
            ],
        },
    },
]

WSGI_APPLICATION = "wsgi.application"
ASGI_APPLICATION = "asgi.application"
CHANNEL_LAYERS = {
    "default": {"BACKEND": "channels.layers.InMemoryChannelLayer"},
}

# Database
# https://docs.djangoproject.com/en/1.11/ref/settings/#databases

DATABASES = {
    "default": {
        "ENGINE": "django.db.backends.sqlite3",
        "NAME": os.path.join(BASE_DIR, "db.sqlite3"),
    }
}


# Password validation
# https://docs.djangoproject.com/en/1.11/ref/settings/#auth-password-validators

AUTH_PASSWORD_VALIDATORS = [
    {
        "NAME": "django.contrib.auth.password_validation.UserAttributeSimilarityValidator",
    },
    {
        "NAME": "django.contrib.auth.password_validation.MinimumLengthValidator",
    },
    {
        "NAME": "django.contrib.auth.password_validation.CommonPasswordValidator",
    },
    {
        "NAME": "django.contrib.auth.password_validation.NumericPasswordValidator",
    },
]


# Internationalization
# https://docs.djangoproject.com/en/1.11/topics/i18n/

LANGUAGE_CODE = "zh-hans"
LANGUAGE_COOKIE_NAME = "blueking_language"
LANGUAGES = [
    ("zh-hans", _("中文")),
    ("en", _("English")),
]
LOCALE_PATHS = (os.path.join(BASE_DIR, "locale"),)

TIME_ZONE = "Asia/Shanghai"

USE_I18N = True

USE_L10N = True

USE_TZ = True

# Static files (CSS, JavaScript, Images)
# https://docs.djangoproject.com/en/1.11/howto/static-files/

STATIC_URL = "/static/"
STATICFILES_DIRS = [
    os.path.join(BASE_DIR, "backend/web_console/static/"),
    os.path.join(BASE_DIR, "backend/static"),
]


REST_FRAMEWORK = {
    "EXCEPTION_HANDLER": "backend.utils.views.custom_exception_handler",
    "DEFAULT_PERMISSION_CLASSES": ("rest_framework.permissions.IsAuthenticated",),
    "DEFAULT_PAGINATION_CLASS": "rest_framework.pagination.LimitOffsetPagination",
    "PAGE_SIZE": 10000,
    "TEST_REQUEST_DEFAULT_FORMAT": "json",
    "DEFAULT_AUTHENTICATION_CLASSES": (
        # 'rest_framework.authentication.BasicAuthentication',
        "rest_framework.authentication.SessionAuthentication",
        # 'rest_framework.authentication.TokenAuthentication',
    ),
    "DEFAULT_RENDERER_CLASSES": (
        # 'backend.utils.renderers.BKAPIRenderer',
        "rest_framework.renderers.JSONRenderer",
        "rest_framework.renderers.BrowsableAPIRenderer",
    ),
    "DATETIME_FORMAT": "%Y-%m-%d %H:%M:%S",
}

# Change default cookie names to avoid conflict
SESSION_COOKIE_NAME = "bcs_sessionid"
CSRF_COOKIE_NAME = "bcs_csrftoken"
# log max bytes：500m
LOG_MAX_BYTES = 500 * 1024 * 1024
# log count: 10
LOG_BACKUP_COUNT = 10


def get_logging_config(log_level, rds_hander_settings=None, log_path="app.log"):
    if not rds_hander_settings:
        rds_hander_settings = {
            "level": "DEBUG",
            "class": "logging.NullHandler",
        }

    rds_hander_settings["filters"] = ["request_id"]

    return {
        "version": 1,
        "disable_existing_loggers": False,
        "formatters": {
            "verbose": {
                "format": "%(levelname)s [%(asctime)s] [%(request_id)s] %(name)s %(pathname)s %(lineno)d %(funcName)s %(process)d %(thread)d \n \t %(message)s \n",  # noqa
                "datefmt": "%Y-%m-%d %H:%M:%S",
            },
            "simple": {"format": "%(levelname)s %(message)s"},
        },
        "filters": {"request_id": {"()": "backend.utils.log.RequestIdFilter"}},
        "handlers": {
            "null": {
                "level": "DEBUG",
                "class": "logging.NullHandler",
            },
            "mail_admins": {"level": "ERROR", "class": "django.utils.log.AdminEmailHandler"},
            "file": {
                "class": "logging.handlers.WatchedFileHandler",
                "level": "DEBUG",
                "formatter": "verbose",
                "filename": log_path,
                "filters": ["request_id"],
            },
            "console": {
                "level": "DEBUG",
                "class": "logging.StreamHandler",
                "formatter": "verbose",
                "filters": ["request_id"],
            },
            "logstash_redis": rds_hander_settings,
            "sentry": {
                "level": "DEBUG",
                "class": "raven.contrib.django.raven_compat.handlers.SentryHandler",
                "dsn": SENTRY_DSN,
            },
        },
        "loggers": {
            "django": {
                "handlers": ["null"],
                "level": "INFO",
                "propagate": True,
            },
            "django.request": {
                "handlers": ["console", "logstash_redis", "file"],
                "level": "ERROR",
                "propagate": True,
            },
            "django.db.backends": {
                "handlers": ["console", "logstash_redis", "file"],
                "level": "INFO",
                "propagate": True,
            },
            "django.security": {"handlers": ["console", "logstash_redis", "file"], "level": "INFO", "propagate": True},
            "root": {
                "handlers": ["console", "logstash_redis", "file"],
                "level": log_level,
                "propagate": False,
            },
            "console": {  # 打印redis日志错误，防止循环错误
                "handlers": ["console", "file"],
                "level": log_level,
                "propagate": False,
            },
            "backend": {
                "handlers": ["console", "logstash_redis", "file"],
                "level": os.getenv("DJANGO_LOG_LEVEL", log_level),
            },
            "requests": {
                "handlers": ["console", "logstash_redis", "file"],
                "level": os.getenv("DJANGO_LOG_LEVEL", "INFO"),
            },
            "bkpaas_auth": {
                "handlers": ["console", "logstash_redis", "file"],
                "level": "DEBUG",
            },
            # 配置iam logger
            'iam': {
                'handlers': ['file'],
                'level': os.getenv('IAM_LOG_LEVEL', 'ERROR'),
                'propagate': False,
            },
            "sentry_logger": {"handlers": ["sentry"], "level": "ERROR"},
        },
    }


STATIC_ROOT = os.path.join(BASE_DIR, "staticfiles")

# request_id的头
REQUEST_ID_HEADER = "HTTP_X_REQUEST_ID"


# ******************************** Helm Config Begin ********************************
HELM_BIN = "/bin/helm"  # helm bin filename
HELM3_BIN = "/bin/helm3"
YTT_BIN = "/bin/ytt"
KUBECTL_BIN = "/bin/kubectl"  # default kubectl bin filename
DASHBOARD_CTL_BIN = "/bin/dashboard-ctl"  # default dashboard ctl filename
KUBECTL_BIN_MAP = {
    "1.8.3": "/bin/kubectl",
    "1.12.3": "/bin/kubectl-v1.12.3",
    "1.14.9": "/bin/kubectl-v1.14.9",
    "1.16.3": "/bin/kubectl-v1.16.3",
    "1.18.12": "/bin/kubectl-v1.18.12",
    "1.20.13": "/bin/kubectl-v1.20.13",
}
KUBECFG = "/root/.kube/config"  # kubectl config path, ex: ~/.kube/config
FORCE_APPLY_CLUSTER_ID = ""  # 强制将资源应用该集群，仅用于开发测试目的, 比如 localkube
KUBECTL_MAX_VISIBLE_LEVEL = 2
HELM_INSECURE_SKIP_TLS_VERIFY = False

DEFAULT_CURATOR_CHART = {
    "name": "chartmuseum-curator",
    "version": "0.9.1",
}

HELM_NEED_REGIST_TO_BKE_WHEN_INIT = False
HELM_HAS_ABILITY_SUPPLY_CHART_REPO_SERVICE = False
# *********************************** Helm Config End *****************************

# 项目地址
DEVOPS_HOST = ""
# 容器服务地址
DEVOPS_BCS_HOST = ""
# 容器服务 API 地址
DEVOPS_BCS_API_URL = ""
# 仓库地址
DEVOPS_ARTIFACTORY_HOST = ""
RUN_ENV = "dev"

BK_KIND_LIST = [1, 2]

SITE_URL = "/"

BK_IAM_APP_URL = ""

# 默认指标数据来源，现在支持bk-data, prometheus
DEFAULT_METRIC_SOURCE = "bk-data"
# 普罗米修斯项目白名单
DEFAULT_METRIC_SOURCE_PROM_WLIST = []

# web_console运行模式, 支持external(平台托管), internal（自己集群托管）
WEB_CONSOLE_MODE = "external"

# web_console kubectld命令
WEB_CONSOLE_KUBECTLD_IMAGE_PATH = ""
WEB_CONSOLE_POD_SPEC = {}
WEB_CONSOLE_PORT = int(os.environ.get("WEB_CONSOLE_PORT", 28800))

# WEB_CONSOLE_MODE 为 external时, 指定的集群ID, token, api_host
WEB_CONSOLE_EXTERNAL_CLUSTER = {"ID": "", "API_TOKEN": "", "API_HOST": ""}

# thanos 查询API
THANOS_HOST = ""
THANOS_AUTH = None

# 客服支持消息
COMMON_CUSTOMER_SUPPORT_MSG = _("联系管理员解决")

# 灰度功能提示消息
GRAYSCALE_FEATURE_MSG = "功能灰度测试中，请联系管理员添加白名单"

# 平台组件部署到的命名空间
BCS_SYSTEM_NAMESPACE = "bcs-system"

# bcs-agent YAML 配置文件模板名
BCS_AGENT_YAML_TEMPLTE_NAME = 'bcs_agent_tmpl.yaml'


# *********************************** Helm Config Begin *****************************
DEFAULT_MANAGE_CLUSTER = {'id': '', 'project_id': ''}

DEFAULT_REPO_NAMESPACE_INFO = {'name': '', 'id': ''}

PLATFORM_REPO_INFO = {'name': 'platform', 'url': '', 'provider': 'chartmuseum', 'project_id': ''}

RGW_CONFIG = {
    "admin_host": "",
    "access_key": "",
    "secret_key": "",
    "admin_endpoint": "",
    "tenant": "",
    "default_policy": "",
    "max_size": 1048576,
}

# 用于区分chart路径
HELM_REPO_ENV = "stag"
PLATFORM_REPO_DOMAIN = ""

HELM_DOC_TRICKS = "https://docs.bk.tencent.com/bcs/Container/helm/Skills.html"
HELM_SYNC_DO_DEPLOY = False
# *********************************** Helm Config End *****************************

SENTRY_DSN = ''

# 默认超级用户
ADMIN_SUPERUSERS = []

# 集群信息变更等的通知人
DEFAULT_OPER_USER = ""

# so初始化错误信息查看
SO_ERROR_MSG = ""

# op系统通知人
OP_MAINTAINERS = []

# 容器服务API测试环境测试用户
DEFAULT_API_TEST_USER = ''

# 提供给流水线API调用的默认用户
PIPELINE_DEFAULT_USER = ""
# 提供给标准运维API调用的默认用户
GCLOUD_DEFAULT_USER = ""

# Mesos中LB的默认仓库域名
DEFAUT_MESOS_LB_JFROG_DOMAIN = ''

# 平台名称
PLAT_SHOW_NAME = "蓝鲸容器管理平台"

# 小游戏示例代码下载链接
RUMPETROLL_DEMO_DOWNLOAD_URL = 'http://bkopen-10032816.file.myqcloud.com/rumpetroll-1.0.0.tgz'

# 直接开启的功能开关，不需要在db中配置
DIRECT_ON_FUNC_CODE = ['HAS_IMAGE_SECRET']

# 默认主键 3.2 版本需要主动指定
DEFAULT_AUTO_FIELD = "django.db.models.AutoField"

# 访问 bcs-api 服务需要的token
BCS_APIGW_TOKEN = os.environ.get("BCS_APIGW_TOKEN", "")
# 访问 bcs-api-gateway 服务需要的AUTHORIZATION
BCS_API_GATEWAY_AUTHORIZATION = os.environ.get("BCS_API_GATEWAY_AUTHORIZATION", "")
# 直连新版bcs api的地址
BCS_APIGW_DOMAIN = {"prod": os.environ.get("BCS_API_GATEWAY_PROD_DOMAIN", "")}


# cluster manager的代理配置
CLUSTER_MANAGER_PROXY = {
    # cluster manager 服务的 host
    "HOST": BCS_APIGW_DOMAIN["prod"],
    # 访问 cluster manager 的 token
    "TOKEN": os.environ.get("BCS_API_GATEWAY_AUTHORIZATION", ""),
    # 前端访问的前缀
    "PREFIX_PATH": "api/cluster_manager/proxy/",
}

# 版本日志放置的路径
CHANGE_LOG_PATH = os.path.join(BASE_DIR, "CHANGELOG")

# 共享集群命名空间的前缀
SHARED_CLUSTER_NS_PREFIX = ""

# API 密钥前端渲染用
BCS_API_HOST = ""

# 基础性能查询数据源
PROM_QUERY_STORE = os.environ.get('BKAPP_PROM_QUERY_STORE', 'BK_MONITOR')

# 蓝鲸监控 unify-query 地址
BK_MONITOR_QUERY_HOST = os.environ.get(
    'BKAPP_BK_MONITOR_QUERY_URL', 'http://bk-monitor-unify-query-http.default.svc.cluster.local:10205'
)

try:
    from .base_ext import *  # noqa
except ImportError as e:
    print(f'Load extension failed: {e}')
