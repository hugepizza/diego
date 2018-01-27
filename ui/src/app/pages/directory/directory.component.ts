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
      // <tr class="unread">
      //   <td>
      //     <div class="checkbox m-t-0 m-b-0">
      //       <input type="checkbox" id="ch1">
      //       <label for="ch1"></label>
      //     </div>
      //   </td>
      //   <td class="hidden-xs"><i class="fa fa-star-o"></i></td>
      //   <td class="hidden-xs">Hritik Roshan</td>
      //   <td class="max-texts">
      //     <a [routerLink]="['/inbox/mail-details']" ></a>
      //     <span class="label label-info m-r-10">Work</span> Lorem ipsum perspiciatis unde omnis iste natus error sit voluptatem
      //   </td>
      //   <td class="hidden-xs"><i class="fa fa-paperclip"></i></td>
      //   <td class="text-right"> 12:30 PM </td>
      // </tr>
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
