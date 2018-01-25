import {Injectable} from '@angular/core';

export interface BadgeItem {
  type: string;
  value: string;
}

export interface ChildrenItems {
  state: string;
  target?: boolean;
  name: string;
  type?: string;
  children?: ChildrenItems[];
}

export interface MainMenuItems {
  state: string;
  main_state?: string;
  target?: boolean;
  name: string;
  type: string;
  icon: string;
  badge?: BadgeItem[];
  children?: ChildrenItems[];
}

export interface Menu {
  label: string;
  main: MainMenuItems[];
}

const MENUITEMS = [
  {
    label: 'Dashboard',
    main: [
      {
        state: 'directory',
        name: '全部文件',
        type: 'link',
        icon: 'icon-list',
      }, {
        state: 'group',
        name: '分类',
        type: 'sub',
        icon: 'icon-grid',
        children: [
          {
            state: 'binary',
            name: '二进制'
          }, {
            state: 'image',
            name: '图片'
          }, {
            state: 'text',
            name: '文本文档'
          }, {
            state: 'video',
            name: '视频'
          }, {
            state: 'others',
            name: '其它'
          }
        ]
      }, {
        state: 'upload-latest',
        name: '最近更新',
        type: 'link',
        icon: 'icon-cloud-upload',
        badge: [
          {
            type: 'primary',
            value: '5'
          }
        ],
      }, {
        state: 'share',
        name: '分享',
        type: 'link',
        icon: 'icon-share',
      },
    ],
  }
];

@Injectable()
export class MenuItems {
  getAll(): Menu[] {
    return MENUITEMS;
  }
}
