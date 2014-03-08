class TerminalStartTabVMItem extends KDCustomHTMLView

  MESSAGE_MAP =
    'started'                : 'Checking VM state'
    'vm is already prepared' : 'READY'

  constructor:(options = {}, data)->

    options.tagName = 'li'
    options.cssClass = KD.utils.curry 'vm-loader-item', options.cssClass

    super options, data

    vm               = @getData()
    { vmController } = KD.singletons

    @loader = new KDLoaderView
      size          : width : 16
      showLoader    : yes
      loaderOptions :
        color       : '#f2f2f2'

    @notice = new KDCustomHTMLView
      tagName : 'i'
      partial : '0%'

    @progress = new KDCustomHTMLView
      tagName  : 'cite'


  handleVMStart:(update)->

    { message, currentStep, totalStep } = update
    if message is 'FINISHED'
      @setClass 'ready'
      @emit 'vm.is.prepared'
      @notice.updatePartial 'READY'
      @loader.hide()
      return

    @unsetClass 'off ready'
    # niceMessage = MESSAGE_MAP[message.toLowerCase()]
    # @notice.updatePartial niceMessage or message

    if message is 'STARTED'
      @loader.show()
      @progress.setCss 'background-color', "#1aaf5d"
      return
    percentage = Math.round currentStep/totalStep*100
    @progress.setWidth percentage, '%'
    @notice.updatePartial "#{percentage}%"


  handleVMStop:(update)->

    { message, currentStep, totalStep } = update
    if message is 'FINISHED'
      @setClass 'off'
      @notice.updatePartial 'OFF'
      @loader.hide()
      return

    @unsetClass 'off ready'
    # niceMessage = MESSAGE_MAP[message.toLowerCase()]
    # @notice.updatePartial niceMessage or message
    if message is 'STARTED'
      @loader.show()
      @progress.setCss 'background-color', "#FF7379"
      @notice.updatePartial "100%"
      return
    percentage = 100 - Math.round currentStep/totalStep*100
    @progress.setWidth percentage, '%'
    @notice.updatePartial "#{percentage}%"


  handleVMInfo:(info)->

    unless info
      @loader.hide()
      @notice.updatePartial 'FAILED'
      return

    { state } = info
    switch state.toLowerCase()
      when "running"
        @notice.updatePartial 'READY'
        @setClass 'ready'
      when "stopped"
        @setClass 'off'
        @notice.updatePartial 'OFF'

    @loader.hide()


  click : ->

    osKite = KD.singletons.vmController.kites[@getData().hostnameAlias]

    if @hasClass 'ready'
      @emit 'VMItemClicked', @getData()
    else if @hasClass 'off'
      @once "vm.is.prepared", => @emit 'VMItemClicked', @getData()
      osKite?.vmOn()
    else
      osKite?.vmOff()


  viewAppended:JView::viewAppended


  pistachio:->
    vm    = @getData()
    alias = vm.hostnameAlias
    """
    <figure>{{> @loader}}</figure>#{alias.replace 'koding.kd.io', 'kd.io'}{{> @notice}}
    {{> @progress}}
    """

