// 单个版本详情
export interface IConfigVersion {
  id: number;
  attachment: {
    app_id: number;
    biz_id: number;
  };
  revision: {
    create_at: string;
    creator: string;
  };
  spec: {
    name: string;
    memo: string;
  };
}

// 单个配置详情
export interface IConfigItem {
  id: number;
  spec: {
    file_mode: string;
    file_type: string;
    memo: string;
    name: string;
    path: string;
    permission: {
      privilege: string;
      user: string;
      user_group: string;
    }
  };
  attachment: {
    biz_id: number;
    app_id: number;
  };
  revision: {
    creator: string;
    create_at: string;
    reviser: string;
    update_at: string;
  };
  file_state: string;
}

// 配置项详情（包含签名信息）
export interface IConfigDetail {
  config_item: IConfigItem;
  content: {
      signature: string;
      byte_size: string;
  }
}

// 文件配置概览内容
export interface IFileConfigContentSummary {
  id?: number;
  name: string;
  signature: string;
  size: string;
  update_at?: string;
}

// 配置对比差异详情
export interface IConfigDiffDetail {
  id: number;
  name: string;
  file_type: string;
  current: {
    signature: string;
    byte_size: string;
    update_at: string;
  }
  base: {
    signature: string;
    byte_size: string;
    update_at: string;
  }
}

// 配置项列表查询接口请求参数
export interface IConfigListQueryParams {
  release_id?: number;
  start?: number;
  limit?: number
}

// 版本列表查询接口请求参数
export interface IConfigVersionQueryParams {
  searchKey?: string;
  start?: number;
  limit?: number;
  all?: boolean;
  deprecated?: boolean;
}