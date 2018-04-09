package org.paw.lang.lexer;


import org.junit.Test;
import org.paw.lang.PawScan;

import java.io.ByteArrayInputStream;
import java.util.ArrayList;
import java.util.List;

public class PawLexerTest {


    String template = ""+
            "{# if target=='HTML' #}  \n" +
            "{% import {HtmlModule} from \"@paw/html\" %}  \n" +
            "{# else #}  \n" +
            "{% import {HtmlModule} from \"@paw/android\" %}  \n" +
            "{# endif #}  \n" +
            "{% import {FormModule} from \"@paw/android\" %}  \n" +
            "{% import {AbViewport} from \"@atomicburst/viewport\" %}  \n" +
            "{% import {AbToolbar} from \"@atomicburst/toolbar\" %}  \n" +
            "{% import {AbButton} from \"@atomicburst/button\" %}  \n" +
            "{% import {AbContainer} from \"@atomicburst/container\" %}  \n" +
            "{% import {AbCombo} from \"@atomicburst/combo\" %}  \n" +
            "{% import {AbTextfield} from \"@atomicburst/textfield\" %}  \n" +
            "{% import {AbDatefield} from \"@atomicburst/datefield\" %}  \n" +
            "{% import {AbCheckbox} from \"@atomicburst/checkbox\" %}  \n" +
            "{% import {AbDatatable} from \"@atomicburst/datatable\" %}  \n" +
            "{% import {AbDatatable_column} from \"@atomicburst/datatable_column\" %}  \n" +
            "{% import {AbPanel} from \"@atomicburst/panel\" %}  \n" +
            "{% import {AbTextareafield} from \"@atomicburst/textareafield\" %}  \n" +
            "{% export CldNfeCteEntradaComponent %} \n" +
            "   {% apply FormModule %}  \n" +
            "   {% apply HtmlModule %}  \n" +
            "   <AbViewport #viewport>\n" +
            "       <AbToolbar\n" +
            "               [AbToolbar]=\"{\n" +
            "                      position:'top',\n" +
            "                      ui:'TOOLS-TOP'\n" +
            "                   }\">\n" +
            "           <div style=\"flex-grow: 1;\"></div>\n" +
            "           <AbButton (click)=\"pesquisar(form)\" [AbButton]=\"{text:'Pesquisar',iconCls:'fa fa-search'}\">\n" +
            "           </AbButton>\n" +
            "           <AbButton (click)=\"buscarNfe(form ,modal1)\" [AbButton]=\"{text:'Busca Nfe SeFaz',iconCls:'fa fa-download'}\">\n" +
            "           </AbButton>\n" +
            "           <AbButton (click)=\"buscarCte(form ,modal1)\" [AbButton]=\"{text:'Busca Cte SeFaz',iconCls:'fa fa-download'}\">\n" +
            "           </AbButton>\n" +
            "           <AbButton (click)=\"downloadALL(form ,modal1)\" [AbButton]=\"{text:'Baixar Todos',iconCls:'fa fa-download'}\">\n" +
            "           </AbButton>\n" +
            "       </AbToolbar>\n" +
            "       <AbContainer [AbLayout_flex]=\"{'orientation':'vertical'}\" style=\"width:100%;height:100%\">\n" +
            "           <form #form=\"ngForm\" >\n" +
            "               <AbContainer [AbLayoutGrid]=\"{units:24}\">\n" +
            "                   <AbCombo #empresaCombo  AbValidationRequired\n" +
            "                             [AbLayoutGridItem]=\"{mm:12,sm:12,md:12}\" #cb=\"AbCombo\"\n" +
            "                             [AbCombo]=\"{\n" +
            "                                 fieldLabel:'Empresa' ,\n" +
            "                                 displayField:'nomeFantasia'\n" +
            "                              }\"\n" +
            "                             [AbStore]=\"{\n" +
            "                               service:empresaService,\n" +
            "                               proxy:{\n" +
            "                                 read:{\n" +
            "                                   method:'listQueryCombo',\n" +
            "                                   extraParams:{\n" +
            "                                       record:contador\n" +
            "                                   }\n" +
            "                                 }\n" +
            "                               },\n" +
            "                               columns:[{\n" +
            "                                   dataIndex:'nomeFantasia',\n" +
            "                                   field:'nomeFantasia'\n" +
            "                               },{\n" +
            "                                   dataIndex:'razaoSocial',\n" +
            "                                   field:'razaoSocial'\n" +
            "                               },{\n" +
            "                                   dataIndex:'cnpj',\n" +
            "                                   field:'cnpj'\n" +
            "                               },{\n" +
            "                                   dataIndex:'empresaId',\n" +
            "                                   field:'empresaId'\n" +
            "                               }]\n" +
            "                             }\"\n" +
            "                             [AbRecord]=\"[pesquisa,'empresaId']\"\n" +
            "                             (next)=\"numeroDocumentoTextfield.focus.emit()\"\n" +
            "                   >\n" +
            "                       <AbTemplate [AbTemplate]=\"{type:'tplItem'}\" let-r=\"record\">\n" +
            "                           <table>\n" +
            "                               <tr>\n" +
            "                                   <th style=\"width: 65%\"></th>\n" +
            "                                   <th style=\"width: 35%\"></th>\n" +
            "                               </tr>\n" +
            "                               <tr>\n" +
            "                                   <td><strong>Razão Social:</strong> {{r.get('nomeFantasia')}}</td>\n" +
            "                                   <td rowspan=\"2\">{{r.get('cnpj')}}</td>\n" +
            "                               </tr>\n" +
            "                               <tr>\n" +
            "                                   <td><strong>Nome Fantasia:</strong> {{r.get('razaoSocial')}}</td>\n" +
            "                               </tr>\n" +
            "                           </table>\n" +
            "                       </AbTemplate>\n" +
            "                   </AbCombo>\n" +
            "                   <AbTextfield #numeroDocumentoTextfield\n" +
            "                                 [AbLayoutGridItem]=\"{mm:12,sm:12,md:2}\"\n" +
            "                                 [AbTextfield]=\"{fieldLabel:'N.Doc'}\"\n" +
            "                                 style=\"width:150px\" [AbRecord]=\"[pesquisa,'numeroDocumento']\"\n" +
            "                                 (next)=\"serieTextfield.focus.emit()\"\n" +
            "                                 AbMask_number>\n" +
            "                   </AbTextfield>\n" +
            "                   <AbTextfield #serieTextfield [AbLayoutGridItem]=\"{mm:12,sm:12,md:2}\"\n" +
            "                                 [AbTextfield]=\"{fieldLabel:'Serie'}\"\n" +
            "                                 style=\"width:150px\" [AbRecord]=\"[pesquisa,'serie']\"\n" +
            "                                 AbMask_number\n" +
            "                                 (next)=\"dataInicioCombo.focus.emit()\"\n" +
            "                   >\n" +
            "                   </AbTextfield>\n" +
            "                   <AbDatefield #dataInicioCombo [AbLayoutGridItem]=\"{mm:12,sm:12,md:4}\"\n" +
            "                                 [AbDatefield]=\"{fieldLabel:'Data Inicio'}\"\n" +
            "                                 [AbRecord]=\"[pesquisa,'dtInicio']\"\n" +
            "                                 (next)=\"dataFinalCombo.focus.emit()\"\n" +
            "                   >\n" +
            "                   </AbDatefield>\n" +
            "                   <AbDatefield #dataFinalCombo [AbLayoutGridItem]=\"{mm:12,sm:12,md:4}\"\n" +
            "                                 [AbDatefield]=\"{fieldLabel:'Data Final'}\"\n" +
            "                                 [AbRecord]=\"[pesquisa,'dtFinal']\"\n" +
            "                                 (next)=\"NFECheckbox.focus.emit()\"\n" +
            "                   >\n" +
            "                   </AbDatefield>\n" +
            "                   <AbCheckbox #NFECheckbox [AbLayoutGridItem]=\"{mm:6,sm:6,md:1}\"\n" +
            "                                (AbRecordChange)=\"$event?1:pesquisa.set('CTE',true)\"\n" +
            "                                [AbCheckbox]=\"{fieldLabel:'CTE'}\"\n" +
            "                                [AbRecord]=\"[pesquisa,'CTE']\"\n" +
            "                                (next)=\"CTECheckbox.focus.emit()\"\n" +
            "                   >\n" +
            "                   </AbCheckbox>\n" +
            "                   <AbCheckbox #CTECheckbox [AbLayoutGridItem]=\"{mm:6,sm:6,md:1}\"\n" +
            "                                (AbRecordChange)=\"$event?1:pesquisa.set('NFCE',true)\"\n" +
            "                                [AbCheckbox]=\"{fieldLabel:'NFE'}\"\n" +
            "                                [AbRecord]=\"[pesquisa,'NFE']\"\n" +
            "                                (next)=\"empresaCombo.focus.emit()\"\n" +
            "                   >\n" +
            "                   </AbCheckbox>\n" +
            "                   <AbCombo #exibirCombo [AbLayoutGridItem]=\"{mm:7,sm:7,md:7}\" #cb=\"AbCombo\"\n" +
            "                             [AbCombo]=\"{\n" +
            "                           fieldLabel:'Exibir',\n" +
            "                           displayField:'nome',\n" +
            "                           keyField:'codigo',\n" +
            "                           localQuery:true\n" +
            "                      }\"\n" +
            "                             (next)=\"empresaCombo.focus.emit()\"\n" +
            "                             [AbStore]=\"{\n" +
            "                      proxy:{\n" +
            "                        read:[{\n" +
            "                             nome:'Todos ',\n" +
            "                             codigo:1\n" +
            "                          },{\n" +
            "                             nome:'Somente Não Manifestado',\n" +
            "                             codigo:6\n" +
            "                          }]\n" +
            "                      }\n" +
            "                  }\"\n" +
            "                             [AbRecord]=\"[pesquisa,'exibir']\">\n" +
            "                   </AbCombo>\n" +
            "               </AbContainer>\n" +
            "           </form>\n" +
            "           <AbDatatable\n" +
            "                   [AbLayout_flex_item]=\"{flex:1}\"\n" +
            "                   #datatableStore=\"AbStore\"\n" +
            "                   [summaryTPL]=\"'Valor Total: \\$${this.currency(this.summary.valorTotal)}'\"\n" +
            "                   #datatableComponent=\"AbDatatable\"\n" +
            "                   (AbStoreSelect)=\"select($event)\"\n" +
            "                   style=\"width: 100%\"\n" +
            "                   [AbDatatable]=\"{pagination:'bottom'}\"\n" +
            "                   [AbStore]=\"{\n" +
            "                   service:xmlService,\n" +
            "                   proxy:{\n" +
            "                     read:{\n" +
            "                       method:'listQueryGrid',\n" +
            "                       extraParams:{\n" +
            "                         pesquisa:pesquisa\n" +
            "                       }\n" +
            "                     }\n" +
            "                   },\n" +
            "                   columnsDisplay:[\n" +
            "                       {\n" +
            "                           dataIndex:'downloadLink',\n" +
            "                           field:'downloadLink',\n" +
            "                           displayField:'downloadLink'\n" +
            "                       }\n" +
            "                   ],\n" +
            "                   columns:[{\n" +
            "                       field:'xmlId',\n" +
            "                       dataIndex:'xmlId'\n" +
            "                   },{\n" +
            "                       field:'l_numeroSaida',\n" +
            "                       dataIndex:'l_numeroSaida'\n" +
            "                   },{\n" +
            "                       field:'data',\n" +
            "                       dataIndex:'data',\n" +
            "                       dateFormat: 'Y-m-d H:i:s.u'\n" +
            "                   },{\n" +
            "                       field:'modelo',\n" +
            "                       dataIndex:'modelo'\n" +
            "                   },{\n" +
            "                       field:'empresaEntradaId.nomeFantasia',\n" +
            "                       dataIndex:'nomeFantasia'\n" +
            "                   },{\n" +
            "                       field:'empresaEntradaId.empresaId',\n" +
            "                       dataIndex:'empresaId'\n" +
            "                   },{\n" +
            "                       field:'empresaEntradaId.cnpj',\n" +
            "                       dataIndex:'empresaCnpj'\n" +
            "                   },{\n" +
            "                       field:'cancelado',\n" +
            "                       dataIndex:'cancelado'\n" +
            "                   },{\n" +
            "                       field:'inutilizado',\n" +
            "                       dataIndex:'inutilizado'\n" +
            "                   },{\n" +
            "                       field:'denegado',\n" +
            "                       dataIndex:'denegado'\n" +
            "                   },{\n" +
            "                       field:'manifestado',\n" +
            "                       dataIndex:'manifestado'\n" +
            "                   },{\n" +
            "                       field:'numeroDocumento',\n" +
            "                       dataIndex:'numeroDocumento'\n" +
            "                   },{\n" +
            "                       field:'serie',\n" +
            "                       dataIndex:'serie'\n" +
            "                   },{\n" +
            "                       field:'valorTotal',\n" +
            "                       dataIndex:'valorTotal'\n" +
            "                   }]\n" +
            "       }\"\n" +
            "           >\n" +
            "               <AbDatatable_column [AbDatatable_column]=\"{\n" +
            "           title:'N.DOC',\n" +
            "           dataIndex:'numeroDocumento',\n" +
            "           minWidth:'60px'\n" +
            "         }\">\n" +
            "               </AbDatatable_column>\n" +
            "               <AbDatatable_column [AbDatatable_column]=\"{\n" +
            "           title:'Serie',\n" +
            "           dataIndex:'serie',\n" +
            "           minWidth:'50px'\n" +
            "         }\">\n" +
            "               </AbDatatable_column>\n" +
            "               <AbDatatable_column [AbDatatable_column]=\"{\n" +
            "           title:'Empresa',\n" +
            "           dataIndex:'nomeFantasia',\n" +
            "           minWidth:'200px'\n" +
            "         }\">\n" +
            "               </AbDatatable_column>\n" +
            "               <AbDatatable_column [AbDatatable_column]=\"{\n" +
            "           title:'Empresa CNPJ',\n" +
            "           dataIndex:'empresaCnpj',\n" +
            "           minWidth:'150px'\n" +
            "         }\">\n" +
            "               </AbDatatable_column>\n" +
            "               <AbDatatable_column [AbDatatable_column]=\"{\n" +
            "           title:'Modelo',\n" +
            "           dataIndex:'modelo',\n" +
            "           minWidth:'60px'\n" +
            "         }\">\n" +
            "               </AbDatatable_column>\n" +
            "               <AbDatatable_column [AbDatatable_column]=\"{\n" +
            "           title:'Data',\n" +
            "           dataIndex:'data',\n" +
            "           dateFormat: 'd/m/Y H:i:s',\n" +
            "           minWidth:'150px'\n" +
            "         }\">\n" +
            "               </AbDatatable_column>\n" +
            "               <AbDatatable_column [AbDatatable_column]=\"{\n" +
            "                   title:'Manifestado',\n" +
            "                   dataIndex:'manifestado',\n" +
            "                   minWidth:'50px'\n" +
            "                 }\">\n" +
            "               </AbDatatable_column>\n" +
            "               <AbDatatable_column [AbDatatable_column]=\"{\n" +
            "           title:'Valor',\n" +
            "           dataIndex:'valorTotal',\n" +
            "           minWidth:'150px',\n" +
            "           currency:true\n" +
            "         }\">\n" +
            "               </AbDatatable_column>\n" +
            "               <AbDatatable_column [AbDatatable_column]=\"{\n" +
            "           title:'Ação',\n" +
            "           dataIndex:'downloadLink',\n" +
            "           minWidth:'100px'\n" +
            "         }\"\n" +
            "                                    (AbDatatable_column_cellClick)=\"download($event,modal1,modalManifestacao)\"\n" +
            "               >\n" +
            "               </AbDatatable_column>\n" +
            "           </AbDatatable>\n" +
            "       </AbContainer>\n" +
            "       <AbTemplate AbModal #modal1=\"AbModal\" [closedOut]=\"false\">\n" +
            "           <AbPanel\n" +
            "                   style=\"\n" +
            "                       height:190px;\n" +
            "                       width: calc( 100% - 200px );\n" +
            "                       margin-left: 100px;\n" +
            "                       margin-top: 100px;\n" +
            "                  \"\n" +
            "                   [AbPanel]=\"{title:'Caregando'}\">\n" +
            "               <AbViewport>\n" +
            "                   <AbToolbar\n" +
            "                           [AbToolbar]=\"{\n" +
            "                      position:'top',\n" +
            "                      ui:'TOOLS-TOP'\n" +
            "                   }\">\n" +
            "                       <div style=\"flex-grow: 1;\"></div>\n" +
            "                       <AbButton [AbButton]=\"{text:'Cancelar'}\" (click)=\"cancelDonwload(modal1)\">\n" +
            "                           <AbTemplate [AbTemplate]=\"{type:'icon'}\">\n" +
            "                               <i class=\"fa fa-close\" aria-hidden=\"true\"></i>\n" +
            "                           </AbTemplate>\n" +
            "                       </AbButton>\n" +
            "                   </AbToolbar>\n" +
            "                   <form #formConcluido=\"ngForm\" style=\"height: 100%;width: 100%\">\n" +
            "                       <AbContainer [AbLayout_flex]=\"{'orientation':'vertical'}\" style=\"width:100%;height:100%\">\n" +
            "                           <progress [attr.max]=\"progressValue\" value=\"10000\">\n" +
            "                               <div class=\"progress-bar\">\n" +
            "                                   <span style=\"width: 80%; \">Progress: 80%</span>\n" +
            "                               </div>\n" +
            "                           </progress>\n" +
            "                           <div style=\"text-align: right;padding-right: 50px;\">{{progressText}}</div>\n" +
            "                       </AbContainer>\n" +
            "                   </form>\n" +
            "               </AbViewport>\n" +
            "           </AbPanel>\n" +
            "       </AbTemplate>\n" +
            "       <AbTemplate AbModal #modalManifestacao=\"AbModal\" [closedOut]=\"false\">\n" +
            "           <AbPanel\n" +
            "                   style=\"\n" +
            "                       height:360px;\n" +
            "                       width: calc( 100% - 200px );\n" +
            "                       margin-left: 100px;\n" +
            "                       margin-top: 100px;\n" +
            "                  \"\n" +
            "                   [AbPanel]=\"{title:'Manifestacao'}\">\n" +
            "               <AbViewport>\n" +
            "                   <AbToolbar\n" +
            "                           [AbToolbar]=\"{\n" +
            "                      position:'top',\n" +
            "                      ui:'TOOLS-TOP'\n" +
            "                   }\">\n" +
            "                       <div style=\"flex-grow: 1;\"></div>\n" +
            "                       <AbButton (click)=\"manifestar(form ,formManifestacao,recordManifestacao,modal1 ,modalManifestacao)\" [AbButton]=\"{text:'Salvar',iconCls:'fa fa-check'}\">\n" +
            "                       </AbButton>\n" +
            "                       <AbButton [AbButton]=\"{text:'Cancelar'}\" (click)=\"cancelManifestacao(modalManifestacao)\">\n" +
            "                           <AbTemplate [AbTemplate]=\"{type:'icon'}\">\n" +
            "                               <i class=\"fa fa-close\" aria-hidden=\"true\"></i>\n" +
            "                           </AbTemplate>\n" +
            "                       </AbButton>\n" +
            "                   </AbToolbar>\n" +
            "                   <form #formManifestacao=\"ngForm\" style=\"height: 100%;width: 100%\">\n" +
            "                       <AbContainer [AbLayout_flex]=\"{'orientation':'vertical'}\" style=\"width:100%;height:100%\">\n" +
            "                           <AbCombo #ManifestacaoTipoField [AbLayoutGridItem]=\"{mm:7,sm:7,md:7}\" #cb=\"AbCombo\"\n" +
            "                                     [AbCombo]=\"{\n" +
            "                                               fieldLabel:'Tipo',\n" +
            "                                               displayField:'nome',\n" +
            "                                               keyField:'codigo',\n" +
            "                                               localQuery:true\n" +
            "                                          }\"\n" +
            "                                         (next)=\"empresaCombo.focus.emit()\"\n" +
            "                                         [AbStore]=\"{\n" +
            "                                         proxy:{\n" +
            "                                           read:[\n" +
            "                                               {\n" +
            "                                                 nome:'Confirmação da Operação ',\n" +
            "                                                 codigo:'210200'\n" +
            "                                               },{\n" +
            "                                                   nome:'Confirmação da Operação ',\n" +
            "                                                   codigo:'210200'\n" +
            "                                               },{\n" +
            "                                                   nome:'Ciencia da Operacao',\n" +
            "                                                   codigo:'210210'\n" +
            "                                               },{\n" +
            "                                                   nome:'Desconhecimento da Operacao',\n" +
            "                                                   codigo:'210220'\n" +
            "                                               },{\n" +
            "                                                   nome:'Operacao nao Realizada',\n" +
            "                                                   codigo:'210240'\n" +
            "                                               }\n" +
            "                                           ]\n" +
            "                                         }\n" +
            "                                      }\"\n" +
            "                                     [AbRecord]=\"[recordManifestacao,'tipoManifestacao']\">\n" +
            "                           </AbCombo>\n" +
            "                           <AbTextareafield\n" +
            "                                   AbValidationRequired\n" +
            "                                   [AbLayoutGridItem]=\"{sm:12,md:12}\"\n" +
            "                                   style=\"height: 200px;\"\n" +
            "                                   [AbTextareafield]=\"{fieldLabel:'Motivo'}\"\n" +
            "                                   [AbRecord]=\"[recordManifestacao,'motivo']\"\n" +
            "                           >\n" +
            "                           </AbTextareafield>\n" +
            "                       </AbContainer>\n" +
            "                   </form>\n" +
            "               </AbViewport>\n" +
            "           </AbPanel>\n" +
            "       </AbTemplate>\n" +
            "   </AbViewport>\n" +
            "{% endexport %}";
   @Test
   public void test1(){
       PawLexer lexer = new PawLexer(new PawScan(template));
       List tokens = new ArrayList();
       while (true){
           PawLexerToken token = lexer.tokinize();
           tokens.add(token);
           if (token.is(PawLexerToken.TOKEN_EOF)){
               break;
           }
       }
       tokens.size();
   }
   @Test
   public void test2(){
       PawScan scan = new PawScan(template);
       int c1 = scan.read();
       int c2 = scan.read();
       int c3 = scan.read();
       scan.addCache(c3);
       scan.addCache(c2);
       scan.addCache(c1);

       c1 = scan.read();
       c2 = scan.read();
       c3 = scan.read();


    }

}
