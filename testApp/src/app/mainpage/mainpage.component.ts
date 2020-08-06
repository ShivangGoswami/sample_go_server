import { Component, OnInit } from '@angular/core';
import { MainServiceService } from '../main-service.service';

@Component({
  selector: 'app-mainpage',
  templateUrl: './mainpage.component.html',
  styleUrls: ['./mainpage.component.scss']
})
export class MainpageComponent implements OnInit {
  greetingMessage;
  constructor(public mainService: MainServiceService) { }

  ngOnInit() {
  }

  getData(){
    this.mainService.getGreeting().subscribe((data)=>{
      this.greetingMessage = data['greeting']
    },(err)=>{
      console.log(err);
    })
  }

}
