import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-directory',
  templateUrl: './directory.component.html',
  styleUrls: ['./directory.component.css']
})
export class DirectoryComponent implements OnInit {
  public fileList = [];

  constructor() {
    this.fileList = [{
      fileName: 'abc.txt',
      fileType: 'text',
      fileSize: 12345,
      fileVersion: '1.2.3',
      fileDesc: 'All right, so I\'m not made of stone.',
      filePath: '/release/test/v1.2.3/abc.txt',
      updated: new Date().getTime(),
      authorName: 'Chuanjian Wang',
      authorId: '123-456',
      stared: true,
      labels: {
        '蓝色':'info',
      }
    }]
  }

  ngOnInit() {
    console.log('hellllo')
  }

  listLabels(lbs: Object) {
    let lbList = [];
    for (let key in lbs) {
      lbList.push({key: key, value: 'label-'+lbs[key]});
    }
    return lbList;
  }
}
