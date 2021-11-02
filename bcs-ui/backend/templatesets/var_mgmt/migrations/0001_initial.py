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
"""
# Generated by Django 1.11.5 on 2018-03-27 03:18
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Variable',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('creator', models.CharField(max_length=32, verbose_name='创建者')),
                ('updator', models.CharField(max_length=32, verbose_name='更新者')),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('updated', models.DateTimeField(auto_now=True)),
                ('is_deleted', models.BooleanField(default=False)),
                ('deleted_time', models.DateTimeField(blank=True, null=True)),
                ('project_id', models.CharField(help_text='0:表示对所有项目生效', max_length=32, verbose_name='项目ID')),
                ('key', models.CharField(max_length=255, verbose_name='KEY')),
                ('name', models.CharField(max_length=255, verbose_name='名称')),
                ('default', models.TextField(help_text="以{'value':'默认值'}格式存储默认值,可以存储字符串和数字", verbose_name='默认值')),
                ('desc', models.TextField(verbose_name='说明')),
                ('category', models.CharField(choices=[('sys', '系统内置'), ('custom', '自定义')], default='custom', max_length=16, verbose_name='类型')),
                ('scope', models.CharField(choices=[('global', '全局变量'), ('cluster', '集群变量'), ('namespace', '命名空间变量')], max_length=16, verbose_name='作用范围')),
            ],
            options={
                'ordering': ('id',),
            },
        ),
    ]