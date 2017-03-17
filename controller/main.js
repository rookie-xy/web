'use strict';

var main = angular.module("main", ['ngRoute', 'controllers', 'services']);
var controllers = angular.module('controllers', ['ngAnimate']);
var services = angular.module('services', []);

// 定义路由
main.config(['$routeProvider', function ($routeProvider) {
    $routeProvider
        .when("/config", {controller: 'configQueryCtrl', templateUrl: 'views/config/list.tpl.html'})
        .otherwise({redirectTo: '/dashboard'});

}]);

// 定义页面标题常量
main.constant('CONSTANT_PAGE_HEADER', {
    '/config': {'title': '配置管理', 'breadcrumbs': ['系统管理', '配置管理']},
});

// 全局方法，设置当前用户、判断是否认证以及设置页面标题和面包屑
main.run(function ($rootScope, $location, CONSTANT_PAGE_HEADER) {
    // 监听路径变化事件，修改页面标题和面包屑
    $rootScope.$on('$locationChangeSuccess', function () {
        var header = CONSTANT_PAGE_HEADER[$location.path()];
        if (header != undefined) {
            $rootScope.pageHeader = header.title;
            $rootScope.breadcrumbs = header.breadcrumbs;
        }
    });
});